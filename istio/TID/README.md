# TID on EKS

## Describe Addon Configuration

```shell
aws eks describe-addon-versions --addon-name tetrate-io_istio-distro
aws eks describe-addon-configuration --addon-name tetrate-io_istio-distro --addon-version v1.20.6-eksbuild.1
```

## Create Addon with configuration

```shell
aws eks create-addon --addon-name tetrate-io_istio-distro --addon-version v1.20.6-eksbuild.1 \
    --cluster-name tidtest-zirain \
    --configuration-values file://istiod-values.yaml 
```

## Cleanup

```shell
aws eks delete-addon --addon-name tetrate-io_istio-distro \
    --region=us-east-2 \
    --cluster tidtest-zirain
```