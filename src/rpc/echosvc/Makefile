.PONEY = all

gencode:
	protoc -I model  model/msg.proto --go_out=plugins=grpc:model

clean:
	@echo "CLEAN"

deployment:
	kubectl apply -f deployments/echosvc-deployment.k8s.yaml ; kubectl apply -f deployments/echosvc-service.k8s.yaml

k8s.restart:
	kubectl rollout restart deployment echosvc

docker: 
	docker build --no-cache -t echosvc:v1.0.0 -f deployments/docker/Dockerfile .