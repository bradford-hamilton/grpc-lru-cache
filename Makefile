GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	bash build/compile_protos.sh

.PHONY: binsize
binsize:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-s -w -extldflags "-static"' -o serverbin . && \
	stat -f%z serverbin && \
	rm serverbin
