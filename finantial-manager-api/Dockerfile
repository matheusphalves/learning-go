FROM golang:1.24

WORKDIR /go/src/app

COPY . .

RUN go mod download && go build -o main cmd/main.go

EXPOSE 8000

CMD ["./main"]