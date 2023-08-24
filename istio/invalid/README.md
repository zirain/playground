# How to use


```shell
for i in {1..1000}; do kubectl apply -f galley_validation_failed.yaml; echo; sleep 1; done
for i in {1..1000}; do kubectl apply -f pilot_total_xds_rejects.yaml; echo; sleep 1; done
```