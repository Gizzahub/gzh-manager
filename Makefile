projectname?=gzh-manager
executablename?=gz

default: help

.PHONY: help
help: ## list makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build golang binary
	@echo "Building $(executablename)..."
	@go build -ldflags "-X main.version=$(shell git describe --always --abbrev=0 --tags)" -o $(executablename)

.PHONY: install
install: build ## install golang binary
#	@go install -ldflags "-X main.version=$(shell git describe --always --abbrev=0 --tags)"
	@echo "Installing $(executablename)..."
	@mv $(executablename) $(shell go env GOPATH)/bin/

.PHONY: run
run: ## run the app
	@go run -ldflags "-X main.version=$(shell git describe --always --abbrev=0 --tags)"  main.go

.PHONY: bootstrap
bootstrap: ## install build deps
	go generate -tags tools tools/tools.go

PHONY: test
test: clean ## display test coverage
	go test --cover -parallel=1 -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | sort -rnk3
	
PHONY: clean
clean: ## clean up environment
	rm -rf coverage.out dist/ $(executablename)
	rm -f $(shell go env GOPATH)/bin/$(executablename)
	rm -f $(shell go env GOPATH)/bin/$(projectname)

PHONY: cover
cover: ## display test coverage
	go test -v -race $(shell go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out
	go tool cover -func=coverage.out

PHONY: fmt
fmt: ## format go files
	gofumpt -w .
	gci write .

PHONY: lint
lint: ## lint go files
	golangci-lint run -c .golang-ci.yml

.PHONY: pre-commit
pre-commit:	## run pre-commit hooks
	pre-commit run --all-files

.PHONY: deploy
deploy:
	# TODO ...
	# $build and deploy
	cp * .$(executablename)
