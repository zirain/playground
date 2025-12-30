# RateLimit


```shell
export GW_IP=$(kubectl get svc eg -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
echo $GW_IP
```


## Multi listeners

```shell
curl http://$GW_IP/echo -HHost:www.example.com -H "x-user-id: two"
curl -ik -v  -H "x-user-id: two" -HHost:www.example.com --resolve "www.example.com:443:$GW_IP" https://www.example.com:443/echo
```