FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/go-prod-stack .
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /app/go-prod-stack  /go-prod-stack
EXPOSE 8080
ENTRYPOINT [ "/go-prod-stack" ]