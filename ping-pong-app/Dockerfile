FROM golang:1.24.4-alpine3.21

WORKDIR /app

COPY . .

RUN go build -o server

EXPOSE 3000

ENV PORT=3000

CMD ["./server"]
