
.PHONY: template
template: 
	@if [ "${name}" != "" ]; then \
		mkdir -p pkg/${name}; \
		touch pkg/${name}/solution.go; \
		touch pkg/${name}/solution_test.go; \
		echo "template ${name}, ok!"; \
	fi
	
.PHONY: rm
rm:
	@if [ "${name}" != "" ]; then \
		rm -rf pkg/${name}; \
	fi

.PHONY: test
test:
	go test $(shell go list ./... | grep '${name}')