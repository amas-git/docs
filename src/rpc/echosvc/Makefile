.PONEY = all
VERSION=1.0.1
gencode:
	protoc -I model  model/msg.proto --go_out=plugins=grpc:model

clean:
	@echo "CLEAN"

deployment:
	kubectl apply -f deployments/echosvc-deployment.k8s.yaml
	kubectl apply -f deployments/echosvc-service.k8s.yaml

k8s.restart:
	kubectl rollout restart deployment echosvc

docker.run:
	 docker stop echosvc ; true
	 docker run -p 8888:8888 --rm --name echosvc -d echosvc:v1.0.0

k8s.test:
	grpcurl -import-path model -proto model/msg.proto  -d '{"id":100, "text":"Hello gRPC"}'  -servername localhost -cacert cert/svc.crt  minikube:30100 model.Echo/say 

local.test: docker.run
	grpcurl -import-path model -proto model/msg.proto  -d '{"id":100, "text":"Hello gRPC"}'  -cacert cert/svc.crt  localhost:8888 model.Echo/say 

docker: 
	docker build --build-arg VERSION=2.0.0 --no-cache -t echosvc:v1.0.0 -f deployments/docker/Dockerfile .

docker.reload: docker docker.run local.test
	@echo reloading...