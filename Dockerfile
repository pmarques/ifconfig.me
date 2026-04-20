ARG BUILDKIT_SBOM_SCAN_CONTEXT=true
FROM golang:1.26.2-alpine3.23@sha256:f85330846cde1e57ca9ec309382da3b8e6ae3ab943d2739500e08c86393a21b1

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
