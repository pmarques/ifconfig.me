# go version of ifconfig.me

This is my firt project using go, a simple HTTP server that returns your public IP address

## Docker

* Build

```bash
docker build -t ifconfig.me .
```

* Run

```bash
docker run -p 8080:80 ifconfig.me
```