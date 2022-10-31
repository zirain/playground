resource "kubernetes_namespace" "istio-system" {
  metadata {
    name = "istio-system"
  }
}
