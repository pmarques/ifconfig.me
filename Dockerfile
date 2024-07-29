ARG BUILDKIT_SBOM_SCAN_CONTEXT=true
FROM golang:1.22.5-alpine3.19@sha256:48aac60d4f50477055967586f60391fb1f3cbdc2e176e36f1f7f3fd0f5380ef7

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
