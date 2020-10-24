GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	bash build/compile_protos.sh

.PHONY: binsize
binsize:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o serverbin . && \
	upx serverbin && \
	stat -f%z serverbin && \
	rm serverbin

.PHONY: docker-build
docker-build:
	docker build -t grpc-lru-cache:latest .

.PHONY: docker-run
docker-run:
	docker run -p 21000:21000 grpc-lru-cache