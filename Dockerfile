FROM ubuntu

WORKDIR /app/

RUN mkdir -p /app/bin
RUN mkdir -p /app/assets

COPY assets/config.toml         assets/config.toml
COPY assets/ip.toml             assets/ip.toml
COPY assets/server_type.toml    assets/server_type.toml
COPY bin/login                  bin/login
COPY bin/gateway                bin/gateway
COPY bin/mgr                    bin/mgr
COPY bin/lobby                  bin/lobby
COPY bin/room                   bin/room

WORKDIR /app/bin
