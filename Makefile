#
#  Copyright Brett Logan. All Rights Reserved.
#
#  SPDX-License-Identifier: Apache-2.0
#

GOTOOLS = counterfeiter
GOTOOLS_BINDIR ?= $(shell go env GOBIN)

go.fqp.counterfeiter := github.com/maxbrunsfeld/counterfeiter/v6

.PHONY: mocks
mocks: tools
	go generate ./...

.PHONY: test
test: unit-tests

.PHONY: tools
tools: $(patsubst %,$(GOTOOLS_BINDIR)/%, $(GOTOOLS))

.PHONY: unit-tests
unit-tests:
	go test ./...

gotool.%:
	$(eval TOOL = ${subst gotool.,,${@}})
	@echo "Building ${go.fqp.${TOOL}} -> $(TOOL)"
	@cd tools && GO111MODULE=on GOBIN=$(abspath $(GOTOOLS_BINDIR)) go install ${go.fqp.${TOOL}}

$(GOTOOLS_BINDIR)/%:
	$(eval TOOL = ${subst $(GOTOOLS_BINDIR)/,,${@}})
	@$(MAKE) gotool.$(TOOL)