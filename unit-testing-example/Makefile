test:
	go test ./... -v --coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html