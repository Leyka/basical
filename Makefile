DIST_DIR=dist
BIN=basical
SRC=cmd/main.go
YML_FILE=config.yml

.PHONY: run build dist clean

run:
	@go run $(SRC)

build:
	@go build -o dist/$(BIN) $(SRC)

dist:
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -o $(DIST_DIR)/"$(BIN)-linux-arm" $(SRC)
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/"$(BIN)-linux-arm64" $(SRC)

clean:
	@rm -rf $(DIST_DIR)
	@go clean
	@go mod tidy
