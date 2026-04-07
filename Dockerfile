FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/server ./server

EXPOSE 8080

ENV PORT=8080

CMD ["./server"]
