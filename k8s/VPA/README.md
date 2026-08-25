# Vertical Pod Autoscaling

## Verify VPA InPlace mode (`InPlaceOrRecreate`)

`InPlaceOrRecreate` is VPA's alpha update mode that resizes a Pod's containers
via the native Kubernetes in-place resize API (no eviction/restart) and only
falls back to evict+recreate when an in-place resize isn't possible.

Requirements:
- Kubernetes 1.33+ (`InPlacePodVerticalScaling` is beta and on by default;
  on older clusters enable it on kube-apiserver + kubelet with
  `--feature-gates=InPlacePodVerticalScaling=true`)
- containerd 1.7.11+ / 1.6.18+, or CRI-O 1.29+
- metrics-server running (the VPA recommender needs it)
- VPA updater and admission-controller started with
  `--feature-gates=InPlaceOrRecreate=true`

### 1. Confirm the cluster feature gate is on

```console
kubectl get --raw /metrics | grep InPlacePodVerticalScaling
# kubernetes_feature_enabled{name="InPlacePodVerticalScaling",stage="BETA"} 1
```

### 2. Install VPA with the alpha feature gate

```console
git clone https://github.com/kubernetes/autoscaler.git
cd autoscaler/vertical-pod-autoscaler
FEATURE_GATES="InPlace=true" ./hack/vpa-up.sh

kubectl get pods -n kube-system -l app=vpa-updater
kubectl get pods -n kube-system -l app=vpa-admission-controller
```

### 3. Deploy the demo workload + VPA

```console
kubectl apply -f k8s/VPA/deployment.yaml
kubectl apply -f k8s/VPA/vpa-inplace.yaml
```

`deployment.yaml` runs [`vish/stress`](https://github.com/vishh/stress) with
`-cpus 1`, i.e. it will happily burn far more CPU than its `100m` request, so
the recommender has an obvious signal to act on. `resizePolicy` on the
container is what allows the kubelet to change cpu/memory without restarting
it — `InPlaceOrRecreate` can't do anything useful without it.

### 4. Watch the resize happen

```console
# pod name should stay the same and RESTARTS should stay 0
kubectl get pods -l app=stress-demo -w
```

```console
# watch the recommendation VPA is computing
kubectl get vpa stress-demo-vpa -w
kubectl describe vpa stress-demo-vpa
```

```console
# confirm the updater actually chose the in-place path
kubectl logs -n kube-system -l app=vpa-updater | grep -i inplace
```

Once the updater acts, inspect the pod's resize status/conditions:

```console
POD=$(kubectl get pod -l app=stress-demo -o jsonpath='{.items[0].metadata.name}')

# desired vs. actually-configured resources
kubectl get pod "$POD" -o jsonpath='{.spec.containers[0].resources}{"\n"}'
kubectl get pod "$POD" -o jsonpath='{.status.containerStatuses[0].resources}{"\n"}'
kubectl get pod "$POD" -o jsonpath='{.status.containerStatuses[0].allocatedResources}{"\n"}'

# resize progress: PodResizeInProgress while applying, PodResizePending
# (reason Deferred/Infeasible) if the node can't satisfy it yet, or absent
# once it's done
kubectl get pod "$POD" -o jsonpath='{.status.conditions}' | jq
```

Success looks like: the Pod keeps its original name, `RESTARTS` never
increments, `status.containerStatuses[0].resources.requests.cpu` climbs above
`100m` to match the VPA recommendation, and the `PodResizeInProgress`
condition disappears once applied. If the updater instead evicts the Pod
(new pod name, `Recreate`-style behavior), the in-place path wasn't taken —
check that `resizePolicy` is set and that both the cluster and VPA feature
gates from steps 1–2 are actually enabled.

If the updater logs keep showing something like:

```console
"Can't in-place update pod, pod's controller not found in replica creator map ..."
"In-place resize failed" error="pod not suitable for in-place update ...: not in replicated pods map"
```

it's usually not a feature-gate problem — it's vpa-updater's disruption
safeguard. Before touching any pod it groups live pods by their controller
and requires at least `--min-replicas` (default `2`) of them to be alive,
since disrupting the only replica of a group isn't provably safe.
`deployment.yaml` runs a single replica, so that check drops the pod from
consideration entirely and it's silently deferred forever. `vpa-inplace.yaml`
sets `updatePolicy.minReplicas: 1` to override the default for this VPA
specifically — the alternatives are scaling the Deployment to 2+ replicas,
or passing `--in-place-skip-disruption-budget=true` to vpa-updater
cluster-wide (safe here since `resizePolicy` is `NotRequired` for both
cpu and memory, i.e. the resize is truly non-disruptive).

### 5. Cleanup

```console
kubectl delete -f k8s/VPA/vpa-inplace.yaml
kubectl delete -f k8s/VPA/deployment.yaml
```
