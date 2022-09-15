# Ambient Get Started


## Get Ambient Istioctl

```
wget https://storage.googleapis.com/istio-build/dev/0.0.0-ambient.191fe680b52c1754ee72a06b3e0d3f9d116f2e82/istioctl-0.0.0-ambient.191fe680b52c1754ee72a06b3e0d3f9d116f2e82-linux-amd64.tar.gz
tar -zxvf istioctl-0.0.0-ambient.191fe680b52c1754ee72a06b3e0d3f9d116f2e82-linux-amd64.tar.gz
mv istioctl /usr/local/bin/istioctl-ambient
```

## Setup Kind Cluster

```console
kind create cluster --image=kindest/node:v1.24.0 --config=- <<EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
name: ambient
nodes:
- role: control-plane
- role: worker
- role: worker
EOF
```

## Install MetalLB
```
bash k8s/metallb/install-metallb.sh
```

## Install Ambient

```console
istioctl-ambient install --set profile=ambient -y --set meshConfig.accessLogFile=/dev/stdout
```

## Deploy Application

```console
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/platform/kube/bookinfo.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/networking/bookinfo-gateway.yaml

kubectl apply -f https://raw.githubusercontent.com/linsun/sample-apps/main/sleep/sleep.yaml
kubectl apply -f https://raw.githubusercontent.com/linsun/sample-apps/main/sleep/notsleep.yaml
```

```console
kubectl exec deploy/sleep -- curl -s istio-ingressgateway.istio-system/productpage | head -n1
kubectl exec deploy/sleep -- curl -s productpage:9080/ | head -n1
kubectl exec deploy/notsleep -- curl -s productpage:9080/ | head -n1
```

## Adding your application to the ambient mesh
```console
kubectl label namespace default istio.io/dataplane-mode=ambient
```

## L7 Traffic

```console
kubectl apply -f - <<EOF
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: Gateway
metadata:
 name: productpage
 annotations:
   istio.io/service-account: bookinfo-productpage
spec:
 gatewayClassName: istio-mesh
EOF


kubectl apply -f - <<EOF
apiVersion: gateway.networking.k8s.io/v1alpha2
kind: Gateway
metadata:
 name: reviews
 annotations:
   istio.io/service-account: bookinfo-reviews
spec:
 gatewayClassName: istio-mesh
EOF
```

```console
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/networking/virtual-service-reviews-90-10.yaml
kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/bookinfo/networking/destination-rule-reviews.yaml
```

Verify traffic:
```console
kubectl exec -it deploy/sleep -- sh -c 'for i in $(seq 1 100); do curl -s http://istio-ingressgateway.istio-system/productpage | grep reviews-v.-; done'
```