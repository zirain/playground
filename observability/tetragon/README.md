# Tetragon

Cilium’s new Tetragon component enables powerful realtime, eBPF-based Security Observability and Runtime Enforcement.

## Creat cluster


## Install Tetragon via helm

```console
helm repo add cilium https://helm.cilium.io
helm repo update

helm install tetragon cilium/tetragon -n kube-system
kubectl rollout status -n kube-system ds/tetragon -w
```

## Deploy Demo Application


```console
kubectl create -f https://raw.githubusercontent.com/cilium/cilium/v1.11/examples/minikube/http-sw-app.yaml
```

To start printing events run:
```
kubectl logs -n kube-system ds/tetragon -c export-stdout -f | tetragon observe --namespace default --pod xwing
```


## Tetragon 解析

BPF程序加载代码位于 [pkg/sensors/base](https://github.com/cilium/tetragon/blob/45fdf406975a428108f9f172803f4049ebe9abb5/pkg/sensors/base/base.go#L16)

`*_v53.o` 用于 `kernal >= 5.3`

```console
bash-5.1# ls -l /var/lib/tetragon/
total 9652
-rw-r--r--    1 root     root          8192 May 13 18:12 bpf_alignchecker.o
-rw-r--r--    1 root     root        460968 May 13 18:12 bpf_execve_event.o
-rw-r--r--    1 root     root        483224 May 13 18:12 bpf_execve_event_v53.o
-rw-r--r--    1 root     root        335824 May 13 18:12 bpf_exit.o
-rw-r--r--    1 root     root        332296 May 13 18:12 bpf_fork.o
-rw-r--r--    1 root     root       1671984 May 13 18:12 bpf_generic_kprobe.o
-rw-r--r--    1 root     root       2009952 May 13 18:12 bpf_generic_kprobe_v53.o
-rw-r--r--    1 root     root        446264 May 13 18:12 bpf_generic_retkprobe.o
-rw-r--r--    1 root     root        466112 May 13 18:12 bpf_generic_retkprobe_v53.o
-rw-r--r--    1 root     root       1632400 May 13 18:12 bpf_generic_tracepoint.o
-rw-r--r--    1 root     root       2002816 May 13 18:12 bpf_generic_tracepoint_v53.o
-rw-r--r--    1 root     root          8192 May 13 18:12 bpf_globals.o
-rw-r--r--    1 root     root          7352 May 13 18:12 bpf_lseek.o
drwxrwxrwx    2 root     root          4096 Aug 20 07:38 metadata
```