# go version of ifconfig.me

[ ![Codeship Status for pmarques/ifconfig.me](https://app.codeship.com/projects/c20e8030-a444-0135-0d5c-4a334dfc4b25/status?branch=master)](https://app.codeship.com/projects/255019)
[![Build Status](https://pmarques.semaphoreci.com/badges/ifconfig.me/branches/master.svg)](https://pmarques.semaphoreci.com/projects/ifconfig.me)
[![CircleCI](https://circleci.com/gh/pmarques/ifconfig.me.svg?style=svg)](https://circleci.com/gh/pmarques/ifconfig.me)
[![Build Status](https://travis-ci.org/pmarques/ifconfig.me.svg?branch=master)](https://travis-ci.org/pmarques/ifconfig.me)
![Docker Build Status](https://img.shields.io/docker/build/patrickfmarques/ifconfig.me)

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

## Kubernetes

```
kubectl apply -f k8s-test.yaml
```

### Debug container

Since there is no OS or utils in the container we need to use Ephemeral Containers although it's currenlty in alpha. For
more info check https://kubernetes.io/docs/concepts/workloads/pods/ephemeral-containers/

```
kubectl debug -it ifconfig.me --image=busybox --target=ifconfig.me-9bf8x
```

# References

* [Static Go Binaries](https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07)
* [Multi-stage Builds](https://docs.docker.com/engine/userguide/eng-image/multistage-build/#use-multi-stage-builds)
