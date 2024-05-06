run:
	go run main.go
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/scan-linux-amd64 main.go
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/scan-mac-amd64 main.go
build-win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/scan-win-amd64.exe main.go
build: build-linux build-mac build-win
make-executable:
	chmod +x bin/scan
clean:
	rm -rf bin
install:
	cp bin/scan /usr/local/bin/scan
