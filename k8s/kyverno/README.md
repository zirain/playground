# Kyverno

```
helm repo add kyverno https://kyverno.github.io/kyverno/
helm repo update
helm search repo kyverno -l
helm install kyverno kyverno/kyverno -n kyverno --create-namespace

helm install kyverno-policies kyverno/kyverno-policies -n kyverno
```

```
helm show values kyverno/kyverno > k8s/kyverno/kyverno.values.yaml
helm show values kyverno/kyverno-policies > k8s/kyverno/kyverno-policies.values.yaml
```