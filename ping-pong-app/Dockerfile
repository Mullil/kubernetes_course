FROM golang:1.24.4-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server

EXPOSE 3000

ENV PORT=3000

CMD ["./server"]
