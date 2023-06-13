.PHONY: proto
proto:
	bash build/compile_protos.sh

.PHONY: docker-build
docker-build:
	docker build -t grpc-lru-cache:latest .

.PHONY: docker-run
docker-run:
	docker run -p 21000:21000 -v $(HOME):$(HOME) grpc-lru-cache
