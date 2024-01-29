ARG BUILDKIT_SBOM_SCAN_CONTEXT=true
FROM golang:1.21.6-alpine3.18@sha256:3354c3a94c3cf67cb37eb93a8e9474220b61a196b13c26f1c01715c301b22a69

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
