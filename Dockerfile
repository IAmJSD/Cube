FROM golang:1.13-stretch AS backend-builder
WORKDIR /opt/cube
COPY . .
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o app

FROM alpine AS certs
RUN apk add --no-cache ca-certificates

FROM scratch
WORKDIR /opt/cube
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=backend-builder /opt/cube .
CMD ["./app"]
