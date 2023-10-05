FROM ubuntu:jammy

LABEL org.opencontainers.image.source=https://github.com/n-Arno/go-http-counter

RUN apt-get update && apt-get install curl -y && \
    rm -rf /var/lib/apt/lists/* && \
    curl -sSL -o /go-http-counter https://github.com/n-Arno/go-http-counter/releases/download/v1.0/go-http-counter-linux-amd64 && \
    chmod +x /go-http-counter

EXPOSE 8080

ENTRYPOINT ["/go-http-counter"]
