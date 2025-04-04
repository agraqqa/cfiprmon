FROM golang:1.24.1 AS builder
WORKDIR /src
COPY . .

RUN cd app && \
    CGO_ENABLED=0 go mod download && \
    CGO_ENABLED=0 go build -o /app &&

FROM alpine:latest
COPY --from=builder /app /app
RUN chmod +x /app

ENTRYPOINT ["/app"]
