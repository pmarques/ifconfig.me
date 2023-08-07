FROM golang:1.20.7-alpine

WORKDIR /go/src/github.com/pmarques/ifconfig.me/
COPY go.mod /go/src/github.com/pmarques/ifconfig.me/
RUN go mod download

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=arm64

COPY . /go/src/github.com/pmarques/ifconfig.me/
RUN go build -installsuffix cgo -o ifconfig.me app/main.go

RUN go test -v ./...

FROM scratch
USER 1000:1000
COPY --from=0 /go/src/github.com/pmarques/ifconfig.me .

EXPOSE 80

ENTRYPOINT ["/ifconfig.me"]
