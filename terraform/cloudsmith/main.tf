
module "istio_monitoring_grafana" {
  source = "terraform.cloudsmith.io/tis-containers/istio-monitoring-grafana/tetrate"
  version = "0.0.0-pre"

  gf_url  = "http://localhost:3000"
  gf_auth = "admin:admin" # default admin password for testing
}