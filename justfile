build:
    docker compose build

up:
    docker compose up

down:
    docker compose down

restart:
    docker compose down
    docker compose up

logs:
    docker compose logs

sqlc:
  sqlc generate

rpc:
    protoc --go_out=. --go-grpc_out=. -I api/proto/ api/proto/plugin.proto

run:
    docker compose down
    templ generate
    docker compose build
    docker compose up

