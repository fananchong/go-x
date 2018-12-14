FROM golang:stretch
WORKDIR /app/
COPY assets/config.toml                     assets/config.toml
COPY assets/ip.toml                         assets/ip.toml
COPY assets/server_type.toml                assets/server_type.toml
COPY bin/go-x                               bin/go-x
COPY bin/*.so                               bin
WORKDIR /app/bin
EXPOSE 30000
EXPOSE 7500
