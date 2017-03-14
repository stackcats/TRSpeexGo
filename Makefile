PREFIX?=$(shell pwd)

.PHONY: clean all fmt vet lint build test save
.DEFAULT: all

all: clean fmt vet lint save build test
dist: build
local: dev

APP_NAME=main
GOLINT := $(shell which golint || echo '')
GODEP := $(shell which godep || echo '')
PKGS := $(shell go list ./... | grep -v ^github.com/stackcats/TRSpeexGo/vendor/)

save:
	@echo "+ $@"
	$(if $(GODEP), , \
		$(error Please install godep: go get github.com/tools/godep))
	@$(GODEP) save $(PKGS)

clean:
	@echo "+ $@"
	@rm -rf "${PREFIX}/bin/"
	@rm -rf "${PREFIX}/Godeps/"
	@rm -rf "${PREFIX}/vendor/"
fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v ^vendor/ | tee /dev/stderr)" || \
		(echo >&2 "+ please format Go code with 'gofmt -s'" && false)

vet:
	@echo "+ $@"
	@go vet -tags "${DOCKER_BUILDTAGS}" $(PKGS)


lint:
	@echo "+ $@"
	$(if $(GOLINT), , \
		$(error Please install golint: `go get -u github.com/golang/lint/golint 需翻墙！`))
	@test -z "$$($(GOLINT) ./... 2>&1 | grep -v ^vendor/ | tee /dev/stderr)"

test:
	@echo "+ $@"
	@go test -test.short $(PKGS)

build:
	@echo "+ $@"
	$(if $(GODEP), , \
		$(error Please install godep: go get github.com/tools/godep))
	@mkdir "${PREFIX}/bin"
	rm -rf ${PREFIX}/vendor/github.com/geekypanda
	@${GODEP} go build -v -o ${APP_NAME}
	@mv ${APP_NAME} "${PREFIX}/bin/"
dev:
	@rm -rf "${PREFIX}/bin/"
	@mkdir "${PREFIX}/bin"
	@go build -v -o ${APP_NAME}
	@mv ${APP_NAME} "${PREFIX}/bin/"
restore:
	@echo "+ $@"
	$(if $(GODEP), , \
		$(error Please install godep: go get github.com/tools/godep))
	@$(GODEP) restore -v
