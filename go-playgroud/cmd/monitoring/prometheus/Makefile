RM := rm -rf
BIN := prom-monitor
DOCKER_IMAGE_HUB := ghcr.io/zirain
DOCKER_IMAGE_NAME := prom-monitor
DOCKER_IMAGE_VER := latest
DOCKER_IMAGE := $(DOCKER_IMAGE_HUB)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VER)

.PHONY : clean
clean:
	docker rmi $(DOCKER_IMAGE) || $(RM) ./$(BIN)
	

build:
	go build -o $(BIN) .

docker.build: build
	docker build -t $(DOCKER_IMAGE) .

docker.push: docker.build
	docker push $(DOCKER_IMAGE)
	docker rmi $(DOCKER_IMAGE)
	docker pull $(DOCKER_IMAGE)