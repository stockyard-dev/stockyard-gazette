FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go mod download && CGO_ENABLED=0 go build -o gazette ./cmd/gazette/

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/gazette .
ENV PORT=9250 DATA_DIR=/data
EXPOSE 9250
CMD ["./gazette"]
