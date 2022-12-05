yum -y install kubelet-1.23.7 kubeadm-1.23.7 kubectl-1.23.7
kubeadm init --image-repository registry.aliyuncs.com/google_containers --kubernetes-version=v1.23.7 --pod-network-cidr=10.10.0.0/16



To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

Alternatively, if you are the root user, you can run:

  export KUBECONFIG=/etc/kubernetes/admin.conf

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.30.12:6443 --token b27i62.nqwhxd1gsj7m13qw \
        --discovery-token-ca-cert-hash sha256:336fa5b917e431c042d2c7a838f523ff56397302bb8d571bfdbc9c1bbd9268e8



systemctl enable kubelet


# 升级内核

# 安装kubelet/kubeadm/kubectl

yum -y install kubelet-1.23.7 kubeadm-1.23.7 kubectl-1.23.7

systemctl enable kubelet

# 安装Master

kubeadm init --image-repository registry.aliyuncs.com/google_containers --kubernetes-version=v1.23.7 --pod-network-cidr=10.10.0.0/16