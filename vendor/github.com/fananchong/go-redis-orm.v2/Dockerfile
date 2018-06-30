FROM golang

WORKDIR /app/bin

COPY bin/redis2go /app/bin/

VOLUME /app/input
VOLUME /app/output

ENTRYPOINT ["./redis2go", "--input_dir=/app/input/", "--output_dir=/app/output/"]

