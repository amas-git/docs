gencode:
	protoc -I model  model/msg.proto --go_out=plugins=grpc:model

clean:
	@echo "CLEAN"

build: 
	docker build -t echosvc:v1.0.0 -f deployments/docker/Dockerfile .