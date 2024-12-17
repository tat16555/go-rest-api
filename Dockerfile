FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y gcc
RUN CGO_ENABLED=1 go build -o main .

CMD ["./main"]
