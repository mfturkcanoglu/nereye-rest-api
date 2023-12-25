FROM golang:latest
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /nereye-api

EXPOSE 4000

CMD [ "/nereye-api" ]