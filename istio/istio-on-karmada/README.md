# 创建证书

```
kubectl create namespace istio-system
# 导入证书
kubectl create secret generic cacerts -n istio-system \
    --from-file=primary/ca-cert.pem \
    --from-file=primary/ca-key.pem \
    --from-file=primary/root-cert.pem \
    --from-file=primary/cert-chain.pem
# 分发证书
kubectl apply -f cacerts-pp.yaml

istioctl manifest generate --set profile=external \
    --set values.global.configCluster=true \
    --set values.global.externalIstiod=false | kubectl apply -f -

```
# 创建istio-operator

istioctl operator dump | kubectl apply -f -
# 分发CRD
kubectl apply -f cpp.yaml
kubectl apply -f operator.yaml

kubectl apply -f primary.yaml
kubectl apply -f member2-join.yaml
kubectl apply -f member2.yaml

```

