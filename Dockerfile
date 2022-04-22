#builder
FROM golang:1.17.2 as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o api app/main.go

#runner
FROM debian:11-slim

COPY --from=builder /app/api /app/ 
COPY --from=builder /app/app/config.json /app/

WORKDIR /app

EXPOSE 8080

CMD ["./api"]