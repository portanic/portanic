build:
    docker-compose build

up:
    docker-compose up

down:
    docker-compose down

restart:
    docker-compose down
    docker-compose up

logs:
    docker-compose logs

sqlc:
  sqlc generate
