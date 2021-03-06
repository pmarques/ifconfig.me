FROM golang:1.16

WORKDIR /go/src/github.com/pmarques/ifconfig.me/
COPY . /go/src/github.com/pmarques/ifconfig.me/

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o ifconfig.me app/main.go

RUN go test -v ./...

FROM scratch
COPY --from=0 /go/src/github.com/pmarques/ifconfig.me .
EXPOSE 80
ENTRYPOINT ["/ifconfig.me"]
