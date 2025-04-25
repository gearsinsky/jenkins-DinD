FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go build -o echo-server

EXPOSE 8080
CMD ["./echo-server"]
