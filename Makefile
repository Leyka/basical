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
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $(DIST_DIR)/"$(BIN)-linux-arm64" $(SRC)
	@cp $(YML_FILE) $(DIST_DIR)/$(YML_FILE)
	@tar -czvf $(DIST_DIR)/$(BIN)-linux-arm64.tar.gz -C $(DIST_DIR) $(BIN)-linux-arm64 $(YML_FILE)

clean:
	@rm -rf $(DIST_DIR)
	@go clean
	@go mod tidy
