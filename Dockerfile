FROM ubuntu:18.04

WORKDIR /app/bin


COPY assets/config.toml /app/bin/config.toml
COPY bin/* /app/bin/
