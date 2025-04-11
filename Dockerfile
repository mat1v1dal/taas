FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/api

# 🧠 Asegura que se compile para Linux AMD64
RUN GOOS=linux GOARCH=amd64 go build -o /taas-api

EXPOSE 8080

CMD ["/taas-api"]
