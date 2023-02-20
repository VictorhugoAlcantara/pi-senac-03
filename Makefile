.PHONY: build win-build linux-build mac-build

build: clean win-build linux-build mac-build

win-build:
	@echo "Building Windows"
	@GOOS=windows GOARCH=amd64 go build -o bin/windows/focustime.exe

linux-build:
	@echo "Building Linux"
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/focustime

mac-build:
	@echo "Building Mac"
	@GOOS=darwin GOARCH=amd64 go build -o bin/mac/focustime

.PHONY: clean

clean:
	@echo "Cleaning"
	@rm -rf bin
	@mkdir bin

.PHONY: run

run:
	@echo "Running"
	@go run main.go