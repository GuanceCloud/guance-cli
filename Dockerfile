FROM alpine:latest
COPY guance /usr/local/bin/guance
ENTRYPOINT ["/usr/local/bin/guance"]
