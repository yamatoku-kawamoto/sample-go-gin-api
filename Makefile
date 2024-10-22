.PHONY: go-run , go-build

BUILD_FLSGS =-ldflags="-s -w" -trimpath

go-run:
	@LOCALONLY=true go run ./cmd

go-build:
	@go build $(BUILD_FLSGS) -tags release -o bin/ivs.exe ./cmd

linux-build:
	@GOOS=linux GOARH=amd64 go build $(BUILD_FLSGS) -tags release -o bin/ivs ./cmd
