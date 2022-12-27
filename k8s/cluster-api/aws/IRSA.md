# IRSA(IAM roles for service accounts)

https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html

## 背景

2014 年，AWS Identity and Access Management 使用 OpenID Connect（OIDC）增加了对联合身份验证的支持。此功能允许您通过支持的身份提供商对 AWS API 调用进行身份验证，并获得有效的 OIDC JSON Web 令牌（JWT）。您可以将此令牌传递到 AWS STS AssumeRoleWithWebIdentity API 操作并接收 IAM 临时角色凭证。您可以使用这些凭证与任意 AWS 服务 交互，包括 Amazon S3 和 DynamoDB。

Kubernetes 长期以来将服务账户用作其内部身份系统。Pods 可以使用自动装载的令牌（这是非 OIDC JWT，只有 Kubernetes API 服务器可以验证）进行 Kubernetes API 服务器的身份验证。这些旧服务账户令牌不会过期，轮换签名密钥是一个困难的过程。在 Kubernetes 版本 1.12 中，添加了对新 ProjectedServiceAccountToken 功能的支持。此功能是也包含服务账户身份的 OIDC JSON Web 令牌，并支持可配置的受众。

Amazon EKS 为包含 ProjectedServiceAccountToken JSON Web 令牌的签名密钥的每个集群托管公有 OIDC 发现端点，这样 IAM 等外部系统就可以验证和接收 Kubernetes 颁发的 OIDC 令牌。

## 好处

- 最低权限 – 您可以将 IAM 权限范围限定到服务账户，并且只有使用该服务账户的 pods 可以访问这些权限。此功能还消除了对 kiam 或 kube2iam 等第三方解决方案的需求。

- 凭证隔离 – 容器组（pod）的容器只能检索与该容器所使用服务账户关联的 IAM 角色的凭证。容器永远无法访问其他 pods 中其他容器所使用的凭证。在使用服务账户的 IAM 角色时，pod's 容器还具有分配给 Amazon EKS 节点 IAM 角色 的权限，除非你阻止 pod 访问 Amazon EC2 实例元数据服务（IMDS）。有关更多信息，请参阅限制对分配给工作节点的实例配置文件的访问。

- 可审核性 – 可通过 AWS CloudTrail 进行访问和事件日志记录以帮助确保可追溯性审核。


## 实现

[IRSA](https://github.com/aws/amazon-eks-pod-identity-webhook/blob/master/SELF_HOSTED_SETUP.md)

1. 创建证书

```console
# Generate the keypair
PRIV_KEY="sa-signer.key"
PUB_KEY="sa-signer.key.pub"
PKCS_KEY="sa-signer-pkcs8.pub"
# Generate a key pair
ssh-keygen -t rsa -b 2048 -f $PRIV_KEY -m pem
# convert the SSH pubkey to PKCS8
ssh-keygen -e -m PKCS8 -f $PUB_KEY > $PKCS_KEY
```

2. 创建S3 BUCKET
aws s3api create-bucket --bucket kurator-operator-test


export HOSTNAME=kurator-operator-test.s3.amazonaws.com
export ISSUER_HOSTPATH=kurator-operator-test.s3.amazonaws.com

3. 创建 discovery.json

```console
cat <<EOF > discovery.json
{
    "issuer": "https://$ISSUER_HOSTPATH",
    "jwks_uri": "https://$ISSUER_HOSTPATH/keys.json",
    "authorization_endpoint": "urn:kubernetes:programmatic_authorization",
    "response_types_supported": [
        "id_token"
    ],
    "subject_types_supported": [
        "public"
    ],
    "id_token_signing_alg_values_supported": [
        "RS256"
    ],
    "claims_supported": [
        "sub",
        "iss"
    ]
}
EOF
```

Build self hosted keys generate tool:

```
go build -o self-hosted-keys ./hack/self-hosted/main.go
chmod +x self-hosted-keys && mv self-hosted-keys  /usr/local/bin/
```

```console
self-hosted-keys -key sa-signer-pkcs8.pub  | jq '.keys += [.keys[0]] | .keys[1].kid = ""' > keys.json
```

```console
aws s3 cp --acl public-read ./discovery.json s3://kurator-operator-test/.well-known/openid-configuration
aws s3 cp --acl public-read ./keys.json s3://kurator-operator-test/keys.json

aws s3 cp ./sa-signer.key s3://kurator-operator-test/sa-signer.key
aws s3 cp ./sa-signer.key.pub s3://kurator-operator-test/sa-signer.key.pub
aws s3 cp ./sa-signer-pkcs8.pub s3://kurator-operator-test/sa-signer-pkcs8.pub
```

创建OIDC

CA_THUMBPRINT=$(openssl s_client -connect kurator-operator-test.s3.amazonaws.com:443 -servername kurator-operator-test.s3.amazonaws.com -showcerts < /dev/null 2>/dev/null  |  openssl x509 -in /dev/stdin -sha1 -noout -fingerprint | cut -d '=' -f 2 | tr -d ':')

aws iam create-open-id-connect-provider \
     --url https://kurator-operator-test.s3.amazonaws.com \
     --thumbprint-list $CA_THUMBPRINT \
     --client-id-list sts.amazonaws.com