build-app: 
	go build -o bin/cats_social cmd/cats_social/*.go

run-app: 
	$(MAKE) build-app && ./bin/cats_social