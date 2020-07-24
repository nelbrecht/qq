FROM alpine:latest as builder
RUN apk add --no-cache ca-certificates tzdata
RUN adduser -D -g '' runuser

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ADD qq /qq
USER runuser
ENTRYPOINT ["/qq"]
