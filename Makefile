MAIN_PATH := cmd/main.go
BIN_PATH := bin/golang-lab

run:
	@go run $(MAIN_PATH)

build:
	@go build -o $(BIN_PATH) $(MAIN_PATH)