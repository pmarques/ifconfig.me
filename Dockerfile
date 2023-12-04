ARG BUILDKIT_SBOM_SCAN_CONTEXT=true
FROM golang:1.21.4-alpine3.18@sha256:70afe55365a265f0762257550bc38440e0d6d6b97020d3f8c85328f00200dd8e

WORKDIR /go/src/github.com/pmarques/ifconfig.me/
COPY go.mod /go/src/github.com/pmarques/ifconfig.me/
RUN go mod download

COPY . /go/src/github.com/pmarques/ifconfig.me/
RUN go test -v ./...

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=arm64
RUN go build -installsuffix cgo -o ifconfig.me app/main.go

FROM scratch
USER 1000:1000
COPY --from=0 /go/src/github.com/pmarques/ifconfig.me .

EXPOSE 80

ENTRYPOINT ["/ifconfig.me"]
