
```shell
k create ns zk
k apply -f istio/zookeeper/zookeeper.yaml -n zk



k create ns zk-injected
k label ns zk-injected istio-injection=enabled
k apply -f istio/zookeeper/zookeeper.yaml -n zk-injected



k label ns kafka istio-injection=enabled
```