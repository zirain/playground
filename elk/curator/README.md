
```
helm install --name curator stable/elasticsearch-curator -f es-curator.yaml

helm upgrade curator stable/elasticsearch-curator -f es-curator.yaml

helm del curator --p

```