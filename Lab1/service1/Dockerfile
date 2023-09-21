FROM golang:1.21-alpine

WORKDIR /app

COPY service1/go.mod ./

RUN go mod download

COPY service1/*.go ./

RUN go build -o /service1

EXPOSE 8081

CMD [ "/service1" ]