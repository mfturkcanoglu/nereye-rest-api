FROM golang:1.21.5 as builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o cmd/app

CMD ["cmd/app"]