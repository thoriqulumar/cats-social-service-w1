include .env

build-app:
	go build -o bin/cats_social cmd/cats_social/*.go

run-app: 
	$(MAKE) build-app && ./bin/cats_social

migrate-up:
	migrate -path database/migration/ -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migrate-down:
	migrate -path database/migration/ -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down
