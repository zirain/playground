NUM_WORKERS 	:= 3
CLUSTER_NAME 	:= envoy-gateway

.PHNOY: create-cluster
create-cluster:
	@NUM_WORKERS=$(NUM_WORKERS) k8s/kind/create-cluster.sh

.PHNOY: delete-cluster
delete-cluster:
	@kind delete clusters $(CLUSTER_NAME)

.PHNOY: sync-image
sync-image:
	@k8s/kind/sync-image.sh

.PHNOY: sync-image
sync-eg-image:
	@k8s/kind/sync-eg.image.sh