.PHONY: create-clusters
create-clusters: ## Create clusters
	@scripts/create-clusters.sh

.PHONY: delete-clusters
delete-clusters: ## Delete clusters
	@kind delete clusters primary remote1 remote2

.PHONY: connect-clusters
connect-clusters:
	@scripts/connect-kind-clusters.sh

.PHONY: install-istio
install-istio: ## Install Istio
	@echo "Installing Istio..."
	@scripts/install-istio.sh

.PHONY: install-app
install-app:
	@echo "Installing Applications..."
	@scripts/install-app.sh

.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} \
			/^[.a-zA-Z0-9\/_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } \
			/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)