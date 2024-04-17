# Generate TID template

```bash
helm template tid tetrate-istio -n tetrate-istio-system --version 1.18.1 --repo https://tetratelabs.github.io/istio-helm \
    --set tid-base.defaultRevision="" \
    --set revision=1-18-1 \
    --set global.istioNamespace=tetrate-istio-system > aws/tid/tid.yaml

helm show values tetrate-istio --repo https://tetratelabs.github.io/istio-helm > aws/tid/tid-values.yaml
```
