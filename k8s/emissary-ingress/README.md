# Emissary

# Install via Helm



## init helm repo

```console
helm repo add datawire https://app.getambassador.io
helm repo update
```

## Create Namespace and Install

```console
kubectl create namespace emissary && \
kubectl apply -f https://app.getambassador.io/yaml/emissary/3.0.0/emissary-crds.yaml
kubectl wait --timeout=90s --for=condition=available deployment emissary-apiext -n emissary-system
helm install emissary-ingress --namespace emissary datawire/emissary-ingress && \
kubectl -n emissary wait --for condition=available --timeout=90s deploy -lapp.kubernetes.io/instance=emissary-ingress
```


```console
helm install -n emissary --create-namespace \
     emissary-ingress datawire/emissary-ingress && \
 kubectl rollout status  -n emissary deployment/emissary-ingress -w
```

curl $(k get svc -nemissary emissary-ingress -o jsonpath="{.status.loadBalancer.ingress[0].ip}")/get