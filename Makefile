GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	bash build/compile_protos.sh