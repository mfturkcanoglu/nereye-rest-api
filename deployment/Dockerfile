FROM golang:latest
WORKDIR /

COPY . ./
RUN go mod download

RUN go build -o /nereye-api

EXPOSE 4000

CMD [ "/nereye-api" ]