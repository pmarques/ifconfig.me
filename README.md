# go version of ifconfig.me

[![Publish Container image to DockerHub](https://github.com/pmarques/ifconfig.me/actions/workflows/docker-hub.yml/badge.svg)](https://github.com/pmarques/ifconfig.me/actions/workflows/docker-hub.yml)
[![Build Status](https://pmarques.semaphoreci.com/badges/ifconfig.me/branches/master.svg)](https://pmarques.semaphoreci.com/projects/ifconfig.me)
[![CircleCI](https://circleci.com/gh/pmarques/ifconfig.me.svg?style=svg)](https://circleci.com/gh/pmarques/ifconfig.me)

This was my first project using go, a simple HTTP server that returns your public IP address.
I'm using this project for demos and learning about Docker, Go and experiment some workflows.

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

## Kubernetes

```
kubectl apply -f k8s
```

### Debug container

Since there is no OS or utils in the container we need to use Ephemeral Containers. For
more info check https://kubernetes.io/docs/concepts/workloads/pods/ephemeral-containers/

```
kubectl -n ifconfig debug -it ifconfig-646c744cbc-8nkfn --image=busybox --target=ifconfig
```

# References

* [Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07)
* [Multi-stage Builds](https://docs.docker.com/engine/userguide/eng-image/multistage-build/#use-multi-stage-builds)
