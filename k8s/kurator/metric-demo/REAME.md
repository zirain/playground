
```console
cat <<EOF | kubectl apply -f -
apiVersion: apps.kurator.dev/v1alpha1
kind: Application
metadata:
  name: metric-demo
  namespace: default
spec:
  source:
    gitRepository:
      interval: 3m0s
      ref:
        branch: master
      timeout: 1m0s
      url: https://github.com/zirain/playground
  syncPolicies:
    - destination:
        fleet: quickstart
      kustomization:
        interval: 5m0s
        path: ./k8s/kurator/metric-demo
        prune: true
        timeout: 2m0s
EOF
```