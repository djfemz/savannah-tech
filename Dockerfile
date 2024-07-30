FROM golang:1.20

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
EXPOSE 8082

RUN go test -v ./api/controllers ./api/services ./api/repositories
RUN go build -v -o main .



CMD ["/app/main"]
