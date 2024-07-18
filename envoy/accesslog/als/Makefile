DOCKER_IMAGE ?= ghcr.io/zirain/als:latest
ECHO_TIMESTAMP := date -u +'%Y/%m/%d %H:%M:%S:%N'
GO_BINARY ?= als

build:
	cd src && GOOS=linux go build -ldflags="-extldflags=-static" -o ../.bin/$(GO_BINARY) ./main.go && chmod +x ../.bin/$(GO_BINARY)

docker: build
	docker build -t $(DOCKER_IMAGE) . --build-arg BINARY=$(GO_BINARY)

docker.push: docker
	docker push $(DOCKER_IMAGE)

clean:
	rm -rf .bin