FROM golang:1.16 AS builder
WORKDIR /app/
COPY go.mod go.sum /
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o chatbot .

FROM alpine:latest as certs
RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

FROM scratch
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/chatbot /

EXPOSE 8000

ENTRYPOINT [ "./chatbot" ]
CMD ["http"]