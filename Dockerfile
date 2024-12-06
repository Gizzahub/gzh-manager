FROM alpine:3.21
COPY gzh-manager /usr/bin/gzh-manager
ENTRYPOINT ["/usr/bin/gzh-manager"]