build:
	go build -o bin/api cmd/api/main.go

start:
	./bin/api

docker-build:
	docker build -t conduit -f docker/api/Dockerfile .

docker-start:
	docker compose up -d

docker-stop:
	docker compose down

dev-start:
	./scripts/dev.sh up

dev-stop:
	./scripts/dev.sh down
