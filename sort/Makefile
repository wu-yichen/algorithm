export GOBIN ?= $(PWD)/bin
export PATH := $(GOBIN):$(PATH)
GOIMPORTS := $(GOBIN)/goimports
GOLANGCI_LINT := $(GOBIN)/golangci-lint

test:
	go test -v ./...

sources = $(shell find . -name '*.go' -not -path './vendor/*')
.PHONY: goimports
goimports: tools
	$(GOBIN)/goimports -w $(sources)

.PHONY: lint
lint: tools
	$(GOBIN)/golangci-lint run $(ARGS)

.PHONY: clean
clean:
	rm -f $(GOIMPORTS) $(GOLANGCI_LINT)

$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports

$(GOLANGCI_LINT):
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

tools: $(GOIMPORTS) $(GOLANGCI_LINT)
