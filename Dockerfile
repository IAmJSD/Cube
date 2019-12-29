FROM golang:1.13-stretch AS backend-builder
WORKDIR /var/cube
COPY . .
RUN go get
RUN go build

FROM alpine AS final-build
RUN apk add --no-cache ca-certificates
WORKDIR /var/cube
COPY --from=backend-builder /var/cube .
CMD ./cube
