# Use the official Golang image to create a build artifact.
FROM golang:1.17 as builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux  go build -o app cmd/main.go 
FROM alpine:3.10
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/app /app/app
COPY templates /app/templates
COPY static /app/static
COPY assets /app/assets

CMD ["/app/app"]