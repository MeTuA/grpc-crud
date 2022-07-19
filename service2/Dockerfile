FROM golang:1.17-alpine AS builder
WORKDIR /source
COPY . /source
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o service2 ./.

FROM alpine:3.9
RUN mkdir /app
WORKDIR /app
RUN mkdir migrations
COPY --from=builder /source/migrations /app/migrations
COPY --from=builder /source/service2 /usr/local/bin
RUN chmod a+x /usr/local/bin/service2

ENTRYPOINT [ "service2" ]