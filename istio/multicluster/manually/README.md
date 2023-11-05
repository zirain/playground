# How to install Istio multicluster manually

- Install istio in primary cluster
- Expose istiod service
- Create remote secret with `istioctl x create-remote-secret --server https://172.18.0.3:6443 --name remote > remote.kubeconfig`
  `172.18.0.3` is ip of the docker container
- Install remote istio with discovery address `kubectl get --kubeconfig=${MAIN_KUBECONFIG} svc -nistio-system istio-eastwestgateway  -o jsonpath="{.status.loadBalancer.ingress[0].ip}"`