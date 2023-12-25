FROM golang:1.21.5 as builder

RUN mkdir /nereye
ADD . /nereye
WORKDIR /nereye

RUN go mod download

RUN go build -o /server cmd/app/main.go

EXPOSE 4000

CMD ["/server"]