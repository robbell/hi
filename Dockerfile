FROM golang:1.15.5-alpine AS builder
RUN mkdir /build
ADD app/ /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o app/hi

FROM alpine:latest AS production
COPY --from=builder /build/app .
EXPOSE 80

ENTRYPOINT ["./hi"]
