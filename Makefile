build:
	@go build -o bin/test1

run: build
	@./bin/test1

test:
	@go test -b ./..
