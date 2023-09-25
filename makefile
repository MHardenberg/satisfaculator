BINARY_NAME = bin/satisfaculator.out

build:
	go build -o $(BINARY_NAME) ./cmd/satisfaculator/main.go

run: build
	./$(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)

