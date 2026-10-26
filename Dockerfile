ARG BUILDKIT_SBOM_SCAN_CONTEXT=true
FROM golang:1.27.3-alpine3.23@sha256:754d5200fc3edae76d3b6598a9a3fac7e74981ee6a5413ab955025f185753ef2

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
