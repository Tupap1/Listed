FROM golang:1.24.5

WORKDIR /app

COPY go.* ./

RUN go mod tidy

COPY . .

RUN go install github.com/air-verse/air@latest

RUN go build -o main main.go

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]


