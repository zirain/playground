resource "helm_release" "base" {
  name       = "base"
  repository = "https://istio-release.storage.googleapis.com/charts"
  chart      = "base"
  version    = null #the latest
  namespace  = "istio-system"
}