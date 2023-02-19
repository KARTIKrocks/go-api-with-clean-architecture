lint:
	@echo "Formatting the code"
	go fmt ./...
	@echo "adding missing imports and removing unnecessary imports"
	goimports -w ./**/*.go
	go vet ./...
	golint  ./... | grep -v vendor/ && exit 1 || exit 0

tidy:
	go mod tidy
	go mod vendor

run:
	go run cmd/main.go

test:
	go test -cover -v ./...