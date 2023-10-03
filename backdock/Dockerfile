# BUILD
FROM golang:alpine as builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ./bin/main ./main.go

# SERVE
FROM alpine:latest
COPY --from=builder /app/bin/main /bin/main

CMD ["./bin/main"]
