## tsbc run

Command used to deploy a new SBC cluster

```
tsbc run [flags]
```

### Examples

```
tsbc run --kamailio-pbx-ip 192.168.1.1 --sbc-fqdn sbc.test1.com --rtp-public-ip 1.1.1.1  --host-ip 192.168.10.1
```

### Options

```
      --db-file string                 sqlite file location, file name must end with .db (default: C:\Users\Sakal/.tsbc/sbc.db)
      --docker-log string              docker log file location (default "/var/log/tsbc/docker.log")
  -h, --help                           help for run
      --host-ip string                 the static lan ip address of the docker host
      --kamailio-image string          kamailio docker image name (default "zeljkoiphouse/kamailio:v0.2")
      --kamailio-new-config            generate new config file for Kamailio (default true)
      --kamailio-pbx-ip string         ip address of internal PBX
      --kamailio-pbx-port string       sip port of internal PBX (default "5060")
      --kamailio-rtpeng-port string    rtp engine signalisation port (default "20001")
      --kamailio-sbc-port string       sbc tls port that will be advertised to MS Teams (default "5061")
      --kamailio-sip-dump              enable sip capture for Kamailio
      --kamailio-udp-sip-port string   sbc udp port that will be advertised to internal PBX (default "5060")
      --log-file string                log file location
      --log-level string               log output level (default "info")
      --rtp-image string               rtp engine docker image name (default "zeljkoiphouse/rtpengine:latest")
      --rtp-max-port string            end port for RTP (default "21000")
      --rtp-min-port string            start port for RTP (default "20501")
      --rtp-public-ip string           public ip for RTP transport
      --rtp-signal-port string         port used to communicate with Kamailio (default "20001")
      --sbc-fqdn string                fqdn that Kamailio will advertise
      --staging string                 set staging environment for LetsEncrypt node (default "false")
      --timezone string                set the timezone (default "Europe/Belgrade")
```

### SEE ALSO

* [tsbc](tsbc.md)	 - TSBC connects your local PBX with MS Teams

###### Auto generated by spf13/cobra on 27-Jan-2023