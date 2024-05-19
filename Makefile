clean:
	go mod tidy
	go mod vendor

build: clean
	go build cmd/main.go

run: build
	go run cmd/main.go