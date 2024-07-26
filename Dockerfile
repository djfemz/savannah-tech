FROM golang:1.20

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
EXPOSE 8080

RUN go build -v -o main .

CMD ["/app/main"]
