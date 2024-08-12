FROM golang:1.22.5-alpine AS builder

WORKDIR /app

COPY proto/go.mod proto/go.sum /app/proto/
COPY apps/product/go.mod apps/product/go.sum /app/apps/product/

RUN go work init && \
    go work use /app/proto && \
    go work use /app/apps/product && \
    go mod download

COPY proto /app/proto
COPY apps/product /app/apps/product

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/product /app/apps/product/cmd/main.go

FROM alpine:3.20

WORKDIR /app

COPY apps/product/configs/config.yaml /app/configs/config.yaml
COPY --from=builder /app/product /app/product

EXPOSE 50051

CMD ["/app/product"]
