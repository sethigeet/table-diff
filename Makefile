generate_mock_data:
	@echo "Please run the 'create_mock_data.sql' script!"

run:
	go run .

build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build . -o bin/table-diff.exe
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build . -o bin/table-diff

.PHONY: run build generate_mock_data