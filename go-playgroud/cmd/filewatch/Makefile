DOCKER_IMAGE ?= ghcr.io/zirain/filewatch:latest
ECHO_TIMESTAMP := date -u +'%Y/%m/%d %H:%M:%S:%N'

build:
	GOOS=linux go build -ldflags="-extldflags=-static" -o .bin/filewatch ./main.go && chmod +x .bin/filewatch

docker: build
	docker build -t $(DOCKER_IMAGE) .

docker.push: docker
	docker push $(DOCKER_IMAGE)

deploy: clean
	kubectl delete -f deploy.yaml --ignore-not-found
	kubectl apply -f deploy.yaml
	kubectl wait --timeout=5m deploy/filewatch --for=condition=Available

test: deploy
	kubectl apply -f cm.yaml
	$(ECHO_TIMESTAMP)
	$(eval POD_NAME := $(shell kubectl get po -l app=filewatch --field-selector=status.phase==Running -o jsonpath='{.items[0].metadata.name}'))
	kubectl patch pod $(POD_NAME) --type='json' -p='[{"op": "add", "path": "/metadata/annotations", "value":{"patched":"true"}}]'
	kubectl logs $(POD_NAME) -f

clean:
	kubectl delete -f deploy.yaml --ignore-not-found
	rm -rf .bin/