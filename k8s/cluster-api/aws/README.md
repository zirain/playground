# CAPA (Cluster API on AWS)

## Initialization 

```console
curl -L https://github.com/kubernetes-sigs/cluster-api/releases/download/v1.2.6/clusterctl-linux-amd64 -o clusterctl

chomod +x clusterctl && mv clusterctl /usr/local/bin
clusterctl version

curl -L https://github.com/kubernetes-sigs/cluster-api-provider-aws/releases/download/v2.0.1/clusterawsadm-linux-amd64 -o clusterawsadm

chomod +x clusterawsadm && mv clusterawsadm /usr/local/bin
clusterawsadm version
```


```
export AWS_REGION=us-east-1 # This is used to help encode your environment variables
export AWS_ACCESS_KEY_ID=<your-access-key>
export AWS_SECRET_ACCESS_KEY=<your-secret-access-key>
export AWS_SESSION_TOKEN=<session-token> # If you are using Multi-Factor Auth.

# The clusterawsadm utility takes the credentials that you set as environment
# variables and uses them to create a CloudFormation stack in your AWS account
# with the correct IAM resources.
clusterawsadm bootstrap iam create-cloudformation-stack

# Create the base64 encoded credentials using clusterawsadm.
# This command uses your environment variables and encodes
# them in a value to be stored in a Kubernetes Secret.
export AWS_B64ENCODED_CREDENTIALS=$(clusterawsadm bootstrap credentials encode-as-profile)

# Finally, initialize the management cluster
clusterctl init --infrastructure aws
```

```
export AWS_REGION=us-east-1
export AWS_SSH_KEY_NAME=default
export AWS_CONTROL_PLANE_MACHINE_TYPE=t3.large
export AWS_NODE_MACHINE_TYPE=t3.large

clusterctl generate cluster capi-quickstart \
  --kubernetes-version v1.25.0 \
  --control-plane-machine-count=3 \
  --worker-machine-count=3 \
  > capi-quickstart.yaml

```

Verify kubelet log:
```
journalctl -xefu kubelet
```

clusterctl get kubeconfig quickstart > /root/.kube/capi-quickstart.kubeconfig
kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.24.1/manifests/calico.yaml
kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig get nodes


kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v1.5.4/cert-manager.yaml
kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig apply -f /root/go/src/github.com/zirain/playground/k8s/cluster-api/aws/pod-identity-webhook.yaml

kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig apply -f /root/go/src/github.com/zirain/playground/k8s/cluster-api/aws-lb-controller/v2_4_5_full.yaml
kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig apply -f /root/go/src/github.com/zirain/playground/k8s/cluster-api/aws-lb-controller/v2_4_5_ingclass.yaml

kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig apply -f /root/go/src/github.com/zirain/playground/k8s/cluster-api/aws-lb-controller/echoserver.yaml

kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig apply -f /root/go/src/github.com/zirain/playground/k8s/cluster-api/aws-lb-controller/echoserver-ingress.yaml

kubectl --kubeconfig=/root/.kube/capi-quickstart.kubeconfig  describe ing echoserver