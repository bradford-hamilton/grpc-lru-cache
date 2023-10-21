.PHONY: proto docker-build docker-run

proto:
	bash compile_protos.sh

docker-build:
	docker build -t grpc-lru-cache:latest .

docker-run:
	docker run -p 21000:21000 -v $(HOME):$(HOME) grpc-lru-cache
