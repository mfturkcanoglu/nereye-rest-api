FROM golang:1.21.5 as builder

WORKDIR /app 

COPY . .

RUN go mod download

WORKDIR /cmd/app

RUN go build -o server .

EXPOSE 4000

CMD ["server"]