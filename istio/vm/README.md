# Virtual Machine Installation


```shell
kubectl create namespace sample
kubectl label namespace sample istio-injection=enabled
```


```shell
cat /var/log/istio/istio.log
journalctl -u istio -n 10
journalctl -u istio > journalctl.log

# Make sure that pod id is accesable from VM.
curl httpbin.sample.svc:8000/get

curl localhost:15000/config_dump?include_eds > out/dump.json
```



## Upgrade istio-sidecar

```shell
sudo dpkg -i ./istio-sidecar-1.25.3.deb
# sudo dpkg -i ./istio-sidecar-1.24.2.deb
sudo systemctl daemon-reload 
sudo systemctl restart istio
sudo systemctl status istio
```

## Uninstall

```shell
sudo systemctl stop istio
sudo dpkg -r istio-sidecar
```