FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go build -o echo-server
RUN groupadd -g 996 docker && usermod -aG docker jenkins

EXPOSE 8080
CMD ["./echo-server"]
