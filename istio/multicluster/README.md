
# Multicluster with istio




## Switch context 

```console
export KUBECONFIG='/root/.kube/istio-primary.config'

export KUBECONFIG='/root/.kube/istio-remotes.config'
kubectl config  use-context remote1
```