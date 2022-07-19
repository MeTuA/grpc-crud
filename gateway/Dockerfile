FROM golang:1.17-alpine AS builder
WORKDIR /source
COPY . /source
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o gateway ./.

FROM alpine:3.9
RUN mkdir /app
WORKDIR /app
COPY --from=builder /source/gateway /usr/local/bin
RUN chmod a+x /usr/local/bin/gateway

ENTRYPOINT [ "gateway" ]