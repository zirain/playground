# Kuma

## Install

1. Install Gateway API

    ```console
    kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.0.0/experimental-install.yaml
    ```

1. Install Kuma with Gateway API support

    ```console
    kumactl install control-plane --set "controlPlane.mode=zone" --experimental-gatewayapi | kubectl apply -f -
    ```

1. Create a Gateway
   
    ```consoel
    echo "apiVersion: gateway.networking.k8s.io/v1beta1
    kind: Gateway
    metadata:
      name: kuma
      namespace: default
    spec:
      gatewayClassName: kuma
      listeners:
      - name: proxy
        port: 8080
        protocol: HTTP
    " | kubectl apply -f -
    ```

## Clean

```console
kumactl install control-plane --set "controlPlane.mode=zone" --experimental-gatewayapi | kubectl delete -f -
```