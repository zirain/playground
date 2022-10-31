resource "helm_release" "base" {
  name       = "base"
  repository = "https://istio-release.storage.googleapis.com/charts"
  chart      = "base"
  version    = null #the latest
  namespace  = "istio-system"

  depends_on = [kubernetes_namespace.istio-system]
}

resource "helm_release" "istiod" {
  name       = "istiod"
  repository = "https://istio-release.storage.googleapis.com/charts"
  chart      = "istiod"
  version    = null #the latest
  namespace  = "istio-system"

  set {
    name  = "meshConfig.accessLogFile"
    value = "/dev/stdout"
  }

  set {
    name  = "meshConfig.outboundTrafficPolicy.mode"
    value = "REGISTRY_ONLY"
  }

  depends_on = [helm_release.base]
}
