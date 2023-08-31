module "istio_monitoring_grafana" {
  source  = "terraform.cloudsmith.io/terraform/istio-monitoring-grafana/tetrate"
  version = "v0.0.1"

  gf_url  = "http://localhost:3000"
  gf_auth = ""
}