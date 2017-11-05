# go version of ifconfig.me

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
