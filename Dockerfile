FROM ubuntu:16.04

WORKDIR /app/bin

COPY assets/config.toml         /app/bin/config.toml
COPY assets/ip.toml             /app/bin/ip.toml
COPY assets/server_type.toml    /app/bin/server_type.toml
COPY bin/login                  /app/bin/login
COPY bin/gateway                /app/bin/gateway
COPY bin/hub                    /app/bin/hub
COPY bin/lobby                  /app/bin/lobby
COPY bin/room                   /app/bin/room

