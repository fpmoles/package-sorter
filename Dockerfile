FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o package-sorter .

FROM alpine:3.20

COPY --from=builder /app/package-sorter /usr/local/bin/package-sorter

ENTRYPOINT ["package-sorter"]
