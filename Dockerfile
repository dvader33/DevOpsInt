FROM golang:1.18.3-alpine3.15

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping

EXPOSE 80

CMD [ "/docker-gs-ping" ]