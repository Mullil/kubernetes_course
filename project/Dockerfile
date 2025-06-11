FROM node:24-alpine3.21 AS frontend-builder

WORKDIR /app
COPY frontend ./frontend

RUN cd frontend && npm install && npm run build

FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

COPY . .

COPY --from=frontend-builder /app/frontend/build ./frontend/build

RUN go build -o server

EXPOSE 3000

ENV PORT=3000

CMD ["./server"]
