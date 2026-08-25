# VPA + Istio sidecar

Same idea as `../README.md` (VPA's `InPlace` update mode resizing a running
pod's containers with no restart), but on a pod that also has an
auto-injected Istio sidecar (`istio-proxy`) — to see what VPA does to a
container it doesn't own the definition of.

The punchline: **the app container resizes with zero restarts, the sidecar
does not.** Istio's injection template has no annotation or field for
setting `resizePolicy` on the `istio-proxy` container it adds, so that
container always ends up with no resize policy at all — which defaults to
`RestartContainer` for both cpu and memory. VPA (or anyone) resizing it in
place therefore restarts just that one container. The pod itself is never
evicted/recreated either way — this is a per-container kubelet decision, not
a per-pod one — so the app container and the rest of the pod are unaffected.

## Prerequisites

Everything `../README.md` requires (metrics-server, VPA installed with the
`InPlace`/`InPlaceOrRecreate` feature gates — see its steps 1–2), **plus**
Istio installed with the sidecar-injection webhook live (`istioctl install`,
or the Helm charts — either is fine, this demo only relies on the
`sidecar.istio.io/inject` annotation being honored).

## 1. Deploy

```console
kubectl apply -f k8s/VPA/istio-sidecar/deployment.yaml -n sidecar
kubectl apply -f k8s/VPA/istio-sidecar/vpa.yaml -n sidecar
```

Confirm injection actually happened before going further — you should see
2/2 containers:

```console
kubectl get pods -l app=stress-demo-istio -n sidecar
# NAME                                  READY   STATUS    RESTARTS   AGE
# stress-demo-istio-xxxxxxxxxx-xxxxx    2/2     Running   0          10s

kubectl get pod -l app=stress-demo-istio -n sidecar \
  -o jsonpath='{.items[0].spec.containers[*].name}{"\n"}'
# stress istio-proxy
```

If it's still `1/1`, injection didn't fire — check that the namespace/webhook
is actually configured for injection before continuing; nothing below this
point will demonstrate anything useful without the sidecar present.

## 2. Watch VPA compute a recommendation for both containers

```console
kubectl get vpa stress-demo-istio-vpa -w
kubectl describe vpa stress-demo-istio-vpa
```

You should see recommendations for both `stress` and `istio-proxy` — the
sidecar's own request was deliberately seeded low (`proxyCPU: 10m`,
`proxyMemory: 32Mi` in `deployment.yaml`) so Envoy's baseline footprint alone
exceeds it, without needing to push any real traffic through the mesh.

## 3. Watch the resize — and the restart-count split

```console
POD=$(kubectl get pod -l app=stress-demo-istio -n sidecar -o jsonpath='{.items[0].metadata.name}')

# pod name should stay this same value the whole time
kubectl get pods -l app=stress-demo-istio -n sidecar -w
```

```console
# per-container restart counts
kubectl get pod "$POD" \
  -o jsonpath='{range .status.containerStatuses[*]}{.name}{"\t"}{.restartCount}{"\n"}{end}'
```

Once vpa-updater acts on both containers, expect:

```
stress          0
istio-proxy     1
```

`stress` climbs to its new cpu/memory (check with the same
`status.containerStatuses[].resources`/`allocatedResources` jsonpath as in
`../README.md`) without ever restarting, because it declares
`resizePolicy: NotRequired`. `istio-proxy`'s restart count increments by
exactly one per resize — check `kubectl describe pod "$POD"` and you'll find
it was a plain container restart (new container ID, same pod), not a pod
eviction: the pod's own name, IP, and the `stress` container are untouched
throughout.

## Cleanup

```console
kubectl delete -f k8s/VPA/istio-sidecar/vpa.yaml
kubectl delete -f k8s/VPA/istio-sidecar/deployment.yaml
```
