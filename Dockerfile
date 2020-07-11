# Build a server binary that will run in an alpine image.
FROM golang:alpine AS server_builder

WORKDIR /app/server
ADD . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /server ./

# Build the front-end web app.
FROM node:alpine AS client_builder

COPY --from=server_builder /app/client ./
RUN npm install
RUN npm run build

# Image that we deploy to production.
FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=server_builder /server ./
COPY --from=client_builder /build ./front-end

EXPOSE 6969
CMD ["./server"]
