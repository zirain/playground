
```shell
yq eval -r '.[0].hostStatuses[] | (.address.socketAddress.address + ":" + (.address.socketAddress.portValue | tostring) + " priority=" + ((.priority // 0) | tostring))' ep.out.yaml
```