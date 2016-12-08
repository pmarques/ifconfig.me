FROM golang:1.7.5

WORKDIR /go/src/github.com/pmarques/ifconfig.me/app/

# this will ideally be built by the ONBUILD below ;)
CMD ["go-wrapper", "run"]

COPY . /go/src/github.com/pmarques/ifconfig.me/
RUN go-wrapper install
