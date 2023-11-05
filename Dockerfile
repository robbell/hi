FROM golang:1.15.5-alpine AS builder
RUN mkdir /build
ADD app/ /build
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -o app/hi
ADD app/templates/ /build/app/templates
ADD app/static/index.html /build/app/static/index.html

FROM alpine:latest AS production
COPY --from=builder /build/app /hi
EXPOSE 8080

WORKDIR /hi
ENTRYPOINT ["./hi"]
