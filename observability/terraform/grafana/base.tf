terraform {
  required_providers {
    grafana = {
      source  = "grafana/grafana"
      version = "2.2.0"
    }
  }
}

provider "grafana" {
  url  = var.gf_url
  auth = var.gf_auth
}

resource "grafana_data_source" "prometheus" {
  type       = "prometheus"
  name       = "Prometheus"
  org_id     = 1
  url        = "http://prometheus:9090"
  is_default = true

  json_data_encoded = jsonencode({
    timeInterval = "5s"
  })
}

resource "grafana_folder" "istio_folder" {
  org_id = 1
  title  = "istio"
}