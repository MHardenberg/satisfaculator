BINARY_NAME = bin/satisfaculator.out

build: go build -o {BINARY_NAME}
    ./cmd/satisfaculator/main.go \
    ./internal/recipes/recipe.go \
    ./internal/recipes/product.go

run: 
    make build
    ./bin/satisfaculator

clean:
    go clean
    rm -f {BINARY_NAME}
