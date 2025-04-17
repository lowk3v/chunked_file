build-macos:
	GOOS=darwin GOARCH=arm64 go build -o ./chunk_file main.go && chmod +x chunk_file;
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./chunk_file main.go && chmod +x chunk_file;
