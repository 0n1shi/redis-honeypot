# Beehive Redis

Beehive Redis is a honeypot which is a Redis server written in Golang.

Just a few Redis commands are implemented. (e.g. `COMMAND`, `KEYS`, `GET` and `SET`)

## Development

- Install 

```
$ brew install go
$ brew install docker
$ brew install wireshark
$ curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
$ brew install goreleaser/tap/goreleaser
```

- Start MySQL and Redis

```bash
$ docker compose up -d
```

- Run main.go

```bash
$ air
```
