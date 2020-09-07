FROM alpine@sha256:7df6db5aa61ae9480f52f0b3a06a140ab98d427f86d8d5de0bedab9b8df6b1c0

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

EXPOSE 9005

ENTRYPOINT ["/app/main", "--scheme", "http", "--port", "9005"]