FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
EXPOSE 8080

RUN go test -v ./...

RUN go build -v -o main .

FROM debian:bullseye-slim

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["/app/main"]
