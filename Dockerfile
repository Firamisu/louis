FROM golang:1.25-alpine AS builder
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd

FROM alpine:latest
WORKDIR /
COPY --from=builder /main /main
EXPOSE 8080
ENTRYPOINT ["/main"]
