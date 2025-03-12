FROM golang:1.21-alphine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=linux go build -o bookstore-app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root

COPY --from=builder /app/bookstore-app .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD [ "./bookstore-app" ]