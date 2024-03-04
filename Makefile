BINARY_NAME=github.com/SkullNBones12/book-base

.Phony: help

help: # Show help for each of the Makefile recipes
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done
.Phony: mod
mod: # Initializes the go.mod file
	go mod init ${BINARY_NAME}
.Phony: fmt
fmt: # Formats all files in current and subdirectories
	go fmt ./...
.Phony: run
run: # runs main.go file
	go run main.go
.Phony: tidy
tidy: # Cleans up dependencies to remove compile errors
	go mod tidy
.Phony: build
build: # Builds binaries for Unix/Windows/Darwin
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go
.Phony: clean
clean: # Removes binaries for Unix/Windows/Darwin
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows
.Phony: vet
vet: # Checks source code for suspicious constructs
	go vet
