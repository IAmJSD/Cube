FROM golang:1.13-stretch AS backend-builder
WORKDIR /var/cube
COPY . .
RUN go get
RUN go build

FROM alpine AS certs
RUN apk add --no-cache ca-certificates

FROM scratch AS final-build
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /var/cube
COPY --from=backend-builder /var/cube .
CMD ./cube
