FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go build -o todo-api

EXPOSE 3030

CMD ./todo-api