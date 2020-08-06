# Build a binary in a temporary container.
FROM golang:alpine AS builder

WORKDIR /src
COPY . .

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o codenames

# Run that binary in a lightweight container.
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /src
COPY --from=builder /src/codenames .

EXPOSE 6969
CMD ["./codenames"]
