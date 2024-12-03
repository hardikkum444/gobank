build:
	@go build -o bin/ecom cmd/migrate/main.go

test:
	@go test -v ./..

run: build
	@./bin/ecom

