.PHONY: proto docker-build docker-run

proto:
	buf generate

docker-build:
	docker build -t bradfordhamilton/grpc-lru-cache:latest .

docker-run:
	docker run -p 21000:21000 -v $(HOME):$(HOME) grpc-lru-cache
