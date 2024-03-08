# Custom 

```shell
cfssl gencert -initca ca-csr.json | cfssljson -bare ca
bash cfssl.sh
```

## References

* https://www.baeldung.com/openssl-self-signed-cert
* https://github.com/coreos/docs/blob/master/os/generate-self-signed-certificates.md
* https://smallstep.com/practical-zero-trust/go-grpc-tls
* https://smallstep.com/docs/step-cli/basic-crypto-operations/#requirements
* https://github.com/cloudflare/cfssl
* 