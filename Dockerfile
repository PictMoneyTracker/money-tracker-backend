FROM golang:1.20.3
WORKDIR /go/src/money-tracker-backend
COPY . .
RUN go build -o bin/server main.go
CMD ["./bin/server"]

