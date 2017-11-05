# go version of ifconfig.me

[ ![Codeship Status for pmarques/ifconfig.me](https://app.codeship.com/projects/c20e8030-a444-0135-0d5c-4a334dfc4b25/status?branch=master)](https://app.codeship.com/projects/255019)
[![Build Status](https://semaphoreci.com/api/v1/pmarques/ifconfig-me/branches/master/badge.svg)](https://semaphoreci.com/pmarques/ifconfig-me)
[![CircleCI](https://circleci.com/gh/pmarques/ifconfig.me.svg?style=svg)](https://circleci.com/gh/pmarques/ifconfig.me)

This was my first project using go, a simple HTTP server that returns your public IP address.
I keep using this project for demos and learning about Docker, Go and experiment some workflows.

## Requirements

 * GO >=1.9 (otherwise the steps in this readme will not work)

## Development

### Run
```
go run app/main.go
```

### Test

```
go test ./...
```

## Create binary

### Build / Compile
```
go build -o ifconfig.me app/main.go
```

### Run

```
./ifconfig.me
```

## Docker

* Build

```bash
docker build -t ifconfig.me .
```

* Run

```bash
docker run -p 8080:80 ifconfig.me
```

# References

* [Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07)
* [Multi-stage Builds](https://docs.docker.com/engine/userguide/eng-image/multistage-build/#use-multi-stage-builds)
