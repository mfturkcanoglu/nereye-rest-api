FROM golang:latest
WORKDIR /

COPY . ./
RUN go mod download

RUN go build -o /cmd/app

EXPOSE 4000

CMD [ "/cmd/app" ]