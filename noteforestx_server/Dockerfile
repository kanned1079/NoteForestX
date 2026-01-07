FROM alpine:3.20

WORKDIR /app

COPY bin/traffic-tester-exec /app/noteforestx-server
COPY config/conf.yaml /app/config/conf.yaml

RUN chmod +x /app/noteforestx-server

ENTRYPOINT ["/app/noteforestx-server server"]