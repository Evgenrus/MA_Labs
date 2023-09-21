FROM golang:1.21-alpine

WORKDIR /app

COPY service2/go.mod ./

RUN go mod download

COPY service2/*.go ./

RUN go build -o /service2

EXPOSE 8084

CMD [ "/service2" ]