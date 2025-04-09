# Dockerfile
FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN go build -o /taas-api

EXPOSE 8080

CMD ["/taas-api"]
