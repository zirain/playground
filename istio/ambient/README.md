# Istio Ambient Mode



```shell
HUB=ghr.io/zirain TAG=ambient make docker.pilot docker.proxyv2 docker.install-cni && kind load docker-image ghr.io/zirain/pilot:ambient --name ambient && kind load docker-image ghr.io/zirain/proxyv2:ambient --name ambient && kind load docker-image ghr.io/zirain/install-cni:ambient --name ambient 

out/linux_amd64/istioctl install --kubeconfig /root/.kube/istio-ambient.config  --set hub=ghr.io/zirain --set tag=ambient --set profile=ambient -f /root/go/src/playground/istio/ambient/iop.yaml

```