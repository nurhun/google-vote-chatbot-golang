FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY internal ./internal
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o bot .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ENV PORT=8080
ENV SA_KEY_PATH="./sa-key.json"
COPY sa-key.json ./
COPY --from=builder /app/bot ./
EXPOSE 8080
CMD ["./bot"]