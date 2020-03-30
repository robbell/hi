FROM golang:1.12.5 AS builder
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o app/hi ./...

FROM alpine:latest AS production
COPY --from=builder /build/app .
EXPOSE 80

ENTRYPOINT ["./hi"]