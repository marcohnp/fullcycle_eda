FROM golang:1.21

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev

COPY . .

RUN go mod tidy

CMD ["go", "run", "./cmd/main.go"]