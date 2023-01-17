
# 在生产中大规模自动化 Istio CA 轮换

原文地址：https://mp.weixin.qq.com/s/75paqvd507_ExHHGszB_-Q


## Install Cert-manager

```console
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.10.1/cert-manager.yaml
```

```
cat << EOF | kubectl apply -f -
---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: selfsigned
  namespace: cert-manager
spec:
  selfSigned: {} # 生产环境请勿使用自签名CA
---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: selfsigned-ca
  namespace: cert-manager
spec:
  isCA: true
  duration: 21600h # 900d
  secretName: selfsigned-ca
  commonName: certmanager-ca
  subject:
    organizations:
      - cert-manager
  issuerRef:
    name: selfsigned
    kind: Issuer
    group: cert-manager.io
---
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: selfsigned-ca
spec:
  ca:
    secretName: selfsigned-ca
EOF
```

## 为 Istio 配置中间 CA

```console
kubectl create namespace istio-system
cat << EOF | kubectl apply -f -
---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: cacerts
  namespace: istio-system
spec:
  secretName: cacerts
  duration: 1440h
  renewBefore: 360h
  commonName: istiod.istio-system.svc
  isCA: true
  usages:
    - digital signature
    - key encipherment
    - cert sign
  dnsNames:
    - istiod.istio-system.svc
  issuerRef:
    name: selfsigned-ca
    kind: ClusterIssuer
    group: cert-manager.io
EOF
```

## 安装 Istio

```console
istioctl operator init
cat << EOF | istioctl apply --skip-confirmation -f -
apiVersion: install.istio.io/v1alpha1
kind: IstioOperator
metadata:
  name: cert-manager
  namespace: istio-system
spec:
  components:
    pilot:
      k8s:
        env:
        - name: AUTO_RELOAD_PLUGIN_CERTS
          value: "true"
EOF
```