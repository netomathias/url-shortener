
FROM golang:alpine AS builder

RUN apk --no-cache add git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/bin/app ./cmd/...

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/bin/app .
COPY --from=builder /app/.env .
CMD ["./app"]