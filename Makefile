.PHONY: build test clean lint fmt


default:
	all

all:
	fmt test

test:
	@echo "[INFO] *****************test***********************"
	@go test -v ./...

ben:
	@echo "[INFO] *****************benchmark**********************"
	@go test -benchmem -bench ./...

fmt:
	@echo "[INFO] ***********************formatting****************************"
	@go fmt ./...