FROM golang:latest

WORKDIR /app

COPY . .

ENV GOOGLE_APPLICATION_CREDENTIALS="/app/go-deploy-414616-49ffc912d41d.json"

WORKDIR /app/todo

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]