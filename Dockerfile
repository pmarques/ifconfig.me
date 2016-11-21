FROM golang:1.7.3

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# this will ideally be built by the ONBUILD below ;)
CMD ["go-wrapper", "run"]

COPY app /go/src/app
RUN go-wrapper install
