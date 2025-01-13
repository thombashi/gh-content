VERSION := v0.0.2

EXTENSION_NAME := gh-content
EXTENSION := thombashi/$(EXTENSION_NAME)

BIN_DIR := $(CURDIR)/bin

GOIMPORTS := goimports


GOIMPORTS := $(BIN_DIR)/goimports
$(GOIMPORTS):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install golang.org/x/tools/cmd/goimports@latest

STATICCHECK := $(BIN_DIR)/staticcheck
$(STATICCHECK):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install honnef.co/go/tools/cmd/staticcheck@latest

TESTIFYILINT := $(BIN_DIR)/testifylint
$(TESTIFYILINT):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install github.com/Antonboom/testifylint@latest

.PHONY: build
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(EXTENSION_NAME) .

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: check
check: $(STATICCHECK) $(TESTIFYILINT)
	$(STATICCHECK) ./...
	go vet ./...
	$(TESTIFYILINT) ./...

.PHONY: fmt
fmt: $(GOIMPORTS) $(TESTIFYILINT)
	$(GOIMPORTS) -w .
	gofmt -w -s .
	$(TESTIFYILINT) -fix ./...

.PHONY: help
help: build
	./$(EXTENSION_NAME) --help

.PHONY: uninstall
uninstall:
	-gh extension remove $(EXTENSION)

.PHONY: install
install: build uninstall
	gh extension install .
	gh extension list

.PHONY: update-package
update-package:
	go run update-kpt-package/*

.PHONEY: push-tag
push-tag:
	git push origin $(VERSION)

.PHONY: tag
tag:
	git tag $(VERSION) -m "Release $(VERSION)"

.PHONY: test
test:
	go test -v ./...

.PHONY: run-test
run-test: build
	$(BIN_DIR)/$(EXTENSION_NAME) --repo cli/cli LICENSE
