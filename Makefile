PROGRAM_NAME = table-diff

generate_mock_data:
	@echo "Please run the 'generate_mock_data.py' script and the execute the 'create_mock_data_large.sql' script"

run:
	go run .

build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/$(PROGRAM_NAME).exe .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(PROGRAM_NAME) .

.PHONY: run build generate_mock_data