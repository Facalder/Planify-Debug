run:
	@echo "Have a nice day!"
	@echo "Running golang application..."
	go run ./cmd/app/main.go

tidy:
	@echo "Tidying go module..."
	go run go mod tidy

build:
	@echo "Building go application..."
