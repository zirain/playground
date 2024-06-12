# How to enable header based stateful session

1. Install istio
   ```console
   istioctl install -y --set values.pilot.env.PILOT_ENABLE_PERSISTENT_SESSION_FILTER=true
   ```
1. Install demo application
   ```console
   kubectl label namespace default istio-injection=enabled --overwrite

   kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld.yaml
   kubectl apply -f https://raw.githubusercontent.com/istio/istio/master/samples/helloworld/helloworld-gateway.yaml
   ```
1. Verify traffic, it will random between v1 and v2.
   ```console
   > curl 172.23.255.200/hello
   Hello version: v1, instance: helloworld-v1-7459d7b54b-mdbml
   > curl 172.23.255.200/hello
   Hello version: v2, instance: helloworld-v2-654d97458-glm9s
   ```
1. Update service label with `istio.io/persistent-session-header: x-session-header`
1. Check the configuration
   ```shell
   istioctl pc l istio-ingressgateway-7d89c996bd-kwgmk -n istio-system --port 8080 -oyaml | grep envoy.filters.http.stateful_session -A 3
   istioctl pc r istio-ingressgateway-7d89c996bd-kwgmk.istio-system -oyaml | grep envoy.filters.http.stateful_session -A 10
   ```
2. Verify traffic, the destination should be always same.
   ```shell
   curl -H "x-session-header: MTAuMjQ0LjAuMTQ4OjUwMDA=" 172.23.255.200/hello
   ```
