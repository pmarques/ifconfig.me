FROM golang:1.9.2

WORKDIR /go/src/github.com/pmarques/ifconfig.me/
COPY . /go/src/github.com/pmarques/ifconfig.me/

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o ifconfig.me app/main.go

FROM scratch
COPY --from=0 /go/src/github.com/pmarques/ifconfig.me .
EXPOSE 80
ENTRYPOINT ["/ifconfig.me"]
