Application = tiu
.PHONY: build run debug release clean
build:
	@go build -ldflags="-s -w"
run:
	@./$(Application)

debug: build
	@./$(Application)

release:
	@echo "building for windows"
	@GOOS=windows CGO_ENABLE=0  GOARCH=amd64 go build -ldflags="-s -w"
	@echo "building for linux"
	@GOOS=linux CGO_ENABLE=0  GOARCH=amd64 go build -ldflags="-s -w"

clean:
	rm -rf $(Application)
	rm -rf $(Application).exe

