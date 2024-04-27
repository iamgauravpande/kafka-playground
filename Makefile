BINARY_NAME=kafka-playground
.DEFAULT_GOAL=run
build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 cmd/main/main.go

run: build
#	./bin/${BINARY_NAME}-linux-amd64
 
clean:
	go clean
	rm -rf bin/${BINARY_NAME}-linux-amd64