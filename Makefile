build:
	@go build -o bin/apiProject

run: build
	@./bin/apiProject

test:
	@go test -v ./...