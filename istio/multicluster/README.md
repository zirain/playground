# Istio Multicluster

## Primary remote 


### Same network

```shell
make create-clusters connect-clusters install-istio install-app
```


### Different networks

```shell
ISTIO_NETWORK_MODE=non-flat make create-clusters connect-clusters install-istio install-app
```


## Cleanup

```shell
make clean
```