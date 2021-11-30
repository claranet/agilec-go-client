default: tests

# Run acceptance tests
.PHONY: tests
tests:
	TF_ACC=1 go test ./tests/... -v $(TESTARGS) -timeout 120m