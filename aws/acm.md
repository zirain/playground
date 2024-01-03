# Create a Certificate with ACM Private Certificate Authority

## creating a root certificate authority (CA) in ACM

```shell
export ROOT_CA_ARN=`aws acm-pca create-certificate-authority \
    --certificate-authority-type ROOT \
    --certificate-authority-configuration \
    "KeyAlgorithm=RSA_2048,
    SigningAlgorithm=SHA256WITHRSA,
    Subject={
        Country=CN,
        State=Zhe Jiang,
        Locality=Hang Zhou,
        Organization=ACM PCA Examples,
        OrganizationalUnit=TLS Example,
        CommonName=svc.cluster.local}" \
        --query CertificateAuthorityArn --output text`
```

## self-sign your root CA

1. Retrieving the CSR contents

    ```shell
    ROOT_CA_CSR=`aws acm-pca get-certificate-authority-csr \
        --certificate-authority-arn ${ROOT_CA_ARN} \
        --query Csr --output text | base64`
    ```
1. Sign the CSR using itself as the issuer

    ```shell
    ROOT_CA_CERT_ARN=`aws acm-pca issue-certificate \
        --certificate-authority-arn ${ROOT_CA_ARN} \
        --template-arn arn:aws:acm-pca:::template/RootCACertificate/V1 \
        --signing-algorithm SHA256WITHRSA \
        --validity Value=10,Type=YEARS \
        --csr "${ROOT_CA_CSR}" \
        --query CertificateArn --output text`
    ```

1. Import the signed certificate as the root CA

    ```shell
    ROOT_CA_CERT=`aws acm-pca get-certificate \
        --certificate-arn ${ROOT_CA_CERT_ARN} \
        --certificate-authority-arn ${ROOT_CA_ARN} \
        --query Certificate --output text | base64 `
    ```

1. Import the certificate

    ```shell
    aws acm-pca import-certificate-authority-certificate \
        --certificate-authority-arn $ROOT_CA_ARN \
        --certificate "${ROOT_CA_CERT}"
    ```

1. Grant permissions to the CA to automatically renew the managed certificates it issues

    ```shell
    aws acm-pca create-permission \
        --certificate-authority-arn $ROOT_CA_ARN \
        --actions IssueCertificate GetCertificate ListPermissions \
        --principal acm.amazonaws.com
    ```

1. Optional: Request a managed certificate from ACM using this CA

    ```shell
    export CERTIFICATE_ARN=`aws acm request-certificate \
        --domain-name "*.howto-k8s-http2.svc.cluster.local" \
        --certificate-authority-arn ${ROOT_CA_ARN} \
        --query CertificateArn --output text`
    ```

### Install cert-manager

```shell
helm repo add jetstack https://charts.jetstack.io
helm upgrade -i \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.13.3 \
  --set installCRDs=true
```

```shell
aws iam create-policy \
    --policy-name AWSAcmPcaIssuer \
    --policy-document file://aws/acm-pca-policy.json
```

eksctl delete iamserviceaccount --cluster appmeshtest-zirain --namespace cert-manager --region=us-east-2 --name aws-privateca-issuer

```shell
eksctl create iamserviceaccount \
    --cluster appmeshtest-zirain \
    --namespace cert-manager \
    --region=us-east-2 \
    --name aws-privateca-issuer \
    --attach-policy-arn  arn:aws:iam::354280132276:policy/AWSAcmPcaIssuer \
    --override-existing-serviceaccounts \
    --approve
```

helm repo add awspca https://cert-manager.github.io/aws-privateca-issuer

```shell
helm upgrade -i aws-privateca-issuer awspca/aws-privateca-issuer \
  --namespace cert-manager \
  --create-namespace \
  --set serviceAccount.create=false \
  --set serviceAccount.name=aws-privateca-issuer
```

## Clenup

```shell
aws iam delete-policy --policy-arn "arn:aws:iam::354280132276:policy/AWSAcmPcaIssuer"
aws acm delete-certificate --certificate-arn $CERTIFICATE_ARN
aws acm-pca update-certificate-authority --certificate-authority-arn $ROOT_CA_ARN --status DISABLED
aws acm-pca delete-certificate-authority --certificate-authority-arn $ROOT_CA_ARN
```