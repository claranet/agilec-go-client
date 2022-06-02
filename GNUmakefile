GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: tests

fmt:
	gofmt -w $(GOFMT_FILES)

tests:
	TF_ACC=1 go test ./tests/... -v $(TESTARGS) -timeout 120m

.PHONY: tests fmt