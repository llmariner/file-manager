.PHONY: default
default: test

include common.mk

.PHONY: test
test: go-test-all

.PHONY: lint
lint: go-lint-all git-clean-check

.PHONY: generate
generate: buf-generate-all

.PHONY: build-server
build-server:
	go build -o ./bin/server ./server/cmd/

.PHONY: build-docker-server
build-docker-server:
	docker build --build-arg TARGETARCH=amd64 -t llm-operator/file-manager-server:latest -f build/server/Dockerfile .

.PHONY: build-vstore
build-vstore:
	go build -o ./bin/vstore ./vstore/cmd/

.PHONY: build-docker-vstore
build-docker-vstore:
	docker build --build-arg TARGETARCH=amd64 -t llm-operator/file-manager-vstore:latest -f build/vstore/Dockerfile .
