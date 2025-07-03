FROM golang:1.24.4-alpine3.22

WORKDIR /app

COPY . .

RUN go build -o server

EXPOSE 3001

ENV PORT=3001

CMD ["./server"]