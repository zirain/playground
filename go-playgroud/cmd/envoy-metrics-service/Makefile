clean:
	rm -f envoy-metrics-service

build:
	GOOS=linux go build -o envoy-metrics-service main.go

docker.build: build
	docker build -t ghcr.io/zirain/envoy-metrics-service:latest .

docker.run:
	docker run -p 8080:8080 ghcr.io/zirain/envoy-metrics-service:latest

docker.push: docker.build
	docker push ghcr.io/zirain/envoy-metrics-service:latest

kube.apply:
	kubectl apply -f k8s/envoy-metrics-service.yaml -n monitoring