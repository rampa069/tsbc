package sbc

import (
	"context"
	"fmt"
	"os"

	"github.com/ZeljkoBenovic/tsbc/db"
	"github.com/docker/docker/client"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
)

type Config struct {
	DomainName, Port               string
	LogLevel                       string
	KamailioImage, RTPEngieneImage string
}

type ISBC interface {
	Run()
}

type sbc struct {
	config   Config
	ctx      context.Context
	dockerCl *client.Client
	logger   hclog.Logger
	db       db.IDB
}

func NewSBC(sbcConfig Config) (ISBC, error) {
	// TODO: logger configurability options
	// create new logger instance
	lg := hclog.New(&hclog.LoggerOptions{
		Name:                 "sbc",
		Level:                hclog.LevelFromString(sbcConfig.LogLevel),
		Color:                hclog.AutoColor,
		ColorHeaderAndFields: true,
	})

	lg.Debug("Creating new docker client instance")

	// create new docker client instance
	dCl, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		lg.Error("Could not instantiate new docker client with opts", "err", err)

		return nil, fmt.Errorf("could not create new docker client instance err=%w", err)
	}

	lg.Debug("New docker client instance created")

	// return sbc instance
	return &sbc{
		config:   sbcConfig,
		ctx:      context.Background(),
		logger:   lg,
		dockerCl: dCl,
		db:       db.NewDB(lg),
	}, nil
}

func (s *sbc) close() {
	err := s.dockerCl.Close()
	if err != nil {
		s.logger.Error("Could not close docker client", "err", err)
	}
}

func (s *sbc) Run() {
	// close connections after this function finishes
	defer s.close()

	// create new database and schema if --fresh flag is set
	if viper.GetBool("fresh") {
		if err := s.db.CreateFreshDB(); err != nil {
			os.Exit(1)
		}
	}

	// save sbc configuration information
	if err := s.db.SaveSBCInformation(); err != nil {
		os.Exit(1)
	}

	// create and run containers
	s.createAndRunContainers()
}
