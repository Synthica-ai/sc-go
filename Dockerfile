FROM golang:1.20-alpine AS builder

ENV GOPATH=/go

RUN apk add --no-cache make gcc musl-dev linux-headers git gettext

ADD . /workspace

WORKDIR /workspace

RUN go build -o /app/bin/app ./server/ \
    && cp .env /app

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app /app

WORKDIR /app/bin

CMD ["./app"]
