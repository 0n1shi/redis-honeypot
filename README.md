# Beehive Redis

![](https://img.shields.io/github/license/0n1shi/beehive-redis)
![](https://img.shields.io/github/v/tag/0n1shi/beehive-redis)

Beehive Redis is a honeypot which is a Redis server written in Golang.

Some Redis commands are implemented. (e.g. `COMMAND`, `KEYS`, `GET` and so on)

```bash
beehive-redis help
NAME:
   Redis server of Beehive honeypot series - TCP server which communicate in Redis protocol.

USAGE:
    [global options] command [command options] [arguments...]

COMMANDS:
   version  Show version
   run      Start tcp server
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

## Get started

1. Get a linux binary from [Releases](https://github.com/0n1shi/redis-honeypot/releases)

2. Make a config file like below.

```yaml
port: 6379
mysql:
  host: 127.0.0.1
  user: beehive
  password: beehive
  db: beehive
```

3. start a Redis server.

```bash
beehive-redis run --config config.yaml
2022/02/20 09:17:27 starting Beehive Redis server ...
```

## Development

- Install

```bash
brew install go
brew install docker
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
brew install goreleaser/tap/goreleaser
```

- Start MySQL and Redis

```bash
docker compose up -d
```

- Run main.go

```bash
air
```
