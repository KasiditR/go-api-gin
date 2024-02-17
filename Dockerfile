FROM golang:latest

WORKDIR /app

COPY . .

ENV GOOGLE_APPLICATION_CREDENTIALS="/app/go-api-gin-e3213-646f1b23c89c.json"

WORKDIR /app/todo

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]