.PHONY: proto docker-build

proto:
	buf generate

docker-build:
	docker build -t bradfordhamilton/grpc-lru-cache:latest .
