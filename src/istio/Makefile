docker: cmd/*.go
	docker context use minikube
	docker build -f cmd/Dockerfile  -t hello:v1.0.0 . 

reload: docker
	kubectl rollout restart deployment hello-a 
