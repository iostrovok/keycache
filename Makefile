# test's options
GODEBUG:=GODEBUG=gocacheverify=1

# paths
GOBIN := $(shell pwd)/bin
MMOCKDIR := $(shell pwd)/mmock/

# Include go binaries into path
export PATH := $(PWD)/bin:$(PATH)

ENV := $(GODEBUG) GO111MODULE=on GOBIN=$(GOBIN)

# Defaults...
all: mod test

# teamcity
install: deps mod
	@echo "Environment installed"

deps:
	@echo "======================================================================"
	@echo 'MAKE: deps...'
	@mkdir -p $(GOBIN)
	$(ENV) GO111MODULE=on go get -u -v github.com/golang/mock/mockgen@latest
	$(ENV) GO111MODULE=on go get -u -v github.com/golang/mock@latest
	# go V16+
	#$(ENV) go install  github.com/golang/mock/mockgen@latest
	#$(ENV) go install  github.com/golang/mock@latest

test:
	@echo "======================================================================"
	@echo "Run 'test' for $(PWD)"
	$(SOURCE_PATH) $(GODEBUG) go test -cover -v --check.format=teamcity --check.name=serverhttp. ./...

test-bench:
	@echo "======================================================================"
	@echo "Run 'test' for $(PWD)"
	$(SOURCE_PATH) $(GODEBUG) go test -bench ./...



mod:
	@echo "======================================================================"
	@echo "Run MOD..."
# 	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod verify
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod tidy
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod vendor
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod download
	GO111MODULE=on GONOSUMDB="*" GOPROXY=direct go mod verify

mock-gen:
	mkdir -p $(MMOCKDIR)
	GO111MODULE=on ./bin/mockgen -package mmock github.com/iostrovok/keycache IItem > $(MMOCKDIR)item_mock.go
	GO111MODULE=on ./bin/mockgen -package mmock github.com/iostrovok/keycache IKeyCache > $(MMOCKDIR)/keycache_mock.go
