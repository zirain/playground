## Setup multi mesh with KinD

1. Create cluster

    ```shell
    CLUSTER_NAME=mesh1 ./create-cluster.sh
    CLUSTER_NAME=mesh2 ./create-cluster.sh
    ```
1. Install mesh 

    ```shell
    istioctl install -f iop/mesh1.yaml -y --kubeconfig ~/.kube/kind-config-mesh1
    istioctl install -f iop/mesh2.yaml -y --kubeconfig ~/.kube/kind-config-mesh2
    ```

1. Install sample applications

    ```shell
    kubectl label namespace default istio-injection=enabled --overwrite --kubeconfig ~/.kube/kind-config-mesh1
    kubectl label namespace default istio-injection=enabled --overwrite --kubeconfig ~/.kube/kind-config-mesh2

    kubectl apply -f samples/ --kubeconfig ~/.kube/kind-config-mesh1
    kubectl apply -f samples/ --kubeconfig ~/.kube/kind-config-mesh2
    ```

1. Install grafana and prometheus

    ```shell
    docker-compose up -d
    ```

1. Install prometheus with remote write in each cluster

    ```shell
    helm template prometheus prometheus-community/prometheus --namespace istio-system \
      -f "prometheus/remote1-values.yml" | kubectl apply -f - --kubeconfig ~/.kube/kind-config-mesh1

    helm template prometheus prometheus-community/prometheus --namespace istio-system \
      -f "prometheus/remote2-values.yml" | kubectl apply -f - --kubeconfig ~/.kube/kind-config-mesh2
    ```