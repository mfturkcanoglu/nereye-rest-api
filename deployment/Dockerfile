FROM golang:1.21.5 as builder

ADD . /go/src
WORKDIR /go/src

RUN go mod download

COPY . ./

RUN go build cmd/app/main.go

EXPOSE 4000

COPY --from=builder /go/src/ /app

CMD ["/app"]