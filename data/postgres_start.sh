docker run -d --rm \
    --name guacamole-pg \
    -e POSTGRES_PASSWORD=postgres \
    -p 5432:5432 \
    postgres
