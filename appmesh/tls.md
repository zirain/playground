
# Create the CA

```
cfssl print-defaults config > config.json
cfssl print-defaults csr > csr.json
cat <<EOF > ca-config.json
{
  "signing": {
    "default": {
      "expiry": "8760h"
    },
    "profiles": {
      "client-server": {
        "expiry": "43800h",
        "usages": [
          "signing",
          "key encipherment",
          "server auth",
          "client auth"
        ]
      }
    }
  }
}
EOF
```
```shell
cfssl genkey -initca - <<EOF | cfssljson -bare ca
{
  "CN": "Test CA",
  "key": {
    "algo": "rsa",
    "size": 2048
  },
  "names": [
    {
      "O": "Tetrate"
    }
  ],
  "ca": {
    "expiry": "131400h"
  }
}
EOF
```
```shell
name="ingress.example.com"
cfssl gencert \
  -ca ca.pem \
  -ca-key ca-key.pem \
  -config ca-config.json \
  -profile client-server \
  -hostname "${name}" \
  - \
<<EOF | cfssljson -bare ${name}
{
	"CN": "${name}",
    "hosts": [
      "${name}"
    ],
	"key": {
		"algo": "rsa",
		"size": 2048
	}
}
EOF
```

```shell
kubectl create ns ingress
kubectl create secret generic ingress-tls \
  -n ingress \
  --from-file=key="${name}-key.pem" \
  --from-file=cert="${name}.pem" \
  --from-file=ca=ca.pem
```

```shell
curl -s --resolve ingress.example.com:443:172.18.255.201 https://ingress.example.com/hello
```