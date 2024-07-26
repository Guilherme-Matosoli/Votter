FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN go build -o /main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /main .

EXPOSE 4000

CMD ["./main"]
