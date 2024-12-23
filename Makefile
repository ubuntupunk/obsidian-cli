BINARY_NAME=obsidian-cli

build-all:
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/${BINARY_NAME}
	GOOS=linux GOARCH=amd64 go build -o bin/linux/${BINARY_NAME}
	GOOS=linux GOARCH=arm64 go build -o bin/rpi/${BINARY_NAME}
	GOOS=windows GOARCH=amd64 go build -o bin/windows/${BINARY_NAME}.exe

build-rpi:
	GOOS=linux GOARCH=arm64 go build -o bin/rpi/${BINARY_NAME}

clean-all:
	go clean
	rm bin/darwin/${BINARY_NAME}
	rm bin/linux/${BINARY_NAME}
	rm bin/rpi/${BINARY_NAME}
	rm bin/windows/${BINARY_NAME}.exe

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

install-darwin:
	cp bin/darwin/${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

install-linux:
	cp bin/linux/${BINARY_NAME} /usr/local/bin/${BINARY_NAME}

install-windows:
	cp bin/windows/${BINARY_NAME}.exe /usr/local/bin/${BINARY_NAME}.exe

install-rpi:
	cp bin/rpi/${BINARY_NAME} /usr/local/bin/${BINARY_NAME}
