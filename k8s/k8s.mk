NUM_WORKERS 	:= 3
CLUSTER_NAME 	:= envoy-gateway

.PHNOY: create-cluster
create-cluster:
	@k8s/kind/create-cluster.sh

.PHNOY: delete-cluster
delete-cluster:
	@kind delete clusters $(CLUSTER_NAME)