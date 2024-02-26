.PHONY: build run
build:
	docker run --rm -v `pwd`:/go/src/go-filter -w /go/src/go-filter \
		-e GOPROXY=https://goproxy.cn \
		golang:1.20 \
		go build -v -o libgolang.so -buildmode=c-shared .

up:
	docker-compose up -d

down:
	docker-compose down

test:
	curl localhost:10000/get?sleep=1s
