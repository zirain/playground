#!/bin/sh
cat > base.yaml
exec kubectl kustomize # you can also use "kustomize build ." if you have it installed.