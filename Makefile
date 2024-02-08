build:
	go build -o bin/$(BINARY_NAME) ./cmd/main.go

run:
	go run ./cmd/

clean:
	go clean
	rm -rf bin/$(BINARY_NAME)