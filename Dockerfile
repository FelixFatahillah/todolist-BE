FROM golang:1.20-alpine

WORKDIR /app

COPY . /app

EXPOSE 3030

CMD ["go", "run", "main.go"]
