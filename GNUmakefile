TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
SDK_ONLY_PKGS=$(shell go list ./... | grep -v "/vendor/")

# Testing Auth
export APIC_HOST=https://73.254.132.17:8443
export APIC_USER=admin
export APIC_PASS=password
export APIC_ALLOW_INSECURE=true

default: build

build: fmtcheck
	@echo "go build SDK and install vendor packages"
	@go build ${SDK_ONLY_PKGS}

test: fmtcheck
	@go test -i $(TEST) || exit 1
	@echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

unit: build
	@echo "go test [unit] SDK and vendor packages"
	@go test -tags "unit" $(TEST) || exit 1

integration: build
	@echo "go test [integration] SDK and vendor packages"
	@go test -tags "integration" $(TEST) || exit 1

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./aws"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

.PHONY: build test verbose unit integration vet fmt fmtcheck errcheck test-compile