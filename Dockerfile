FROM ubuntu:16.04

WORKDIR /app/bin


COPY assets/config.toml /app/bin/config.toml
COPY assets/ip.toml /app/bin/ip.toml
COPY bin/login /app/bin/login
COPY bin/gateway /app/bin/gateway

