GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

ifneq ($(origin TESTS), undefined)
	RUNARGS = -run='$(TESTS)'
endif

default: tests

fmt:
	gofmt -w $(GOFMT_FILES)

tests:
	go test  -v ./tests/... $(RUNARGS) -timeout 120m

.PHONY: tests fmt