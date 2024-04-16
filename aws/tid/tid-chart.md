# Generate TID template

```bash
helm template tid tetrate-istio -n tetrate-istio-system --version 1.18.1 --repo https://tetratelabs.github.io/istio-helm --set global.istioNamespace=tetrate-istio-system > aws/tid.yaml

helm show values tetrate-istio --repo https://tetratelabs.github.io/istio-helm > aws/tid-values.yaml
```
