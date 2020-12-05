test-coverage:
	go test -v -coverprofile=coverage.txt -covermode=atomic -timeout=5m -race ./...

test:
	go test -v -timeout=5m -race ./...