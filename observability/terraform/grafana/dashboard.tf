resource "grafana_dashboard" "influxdb" {
  for_each    = fileset("${path.module}/dashboards", "*.json")
  config_json = file("${path.module}/dashboards/${each.key}")
  folder      = grafana_folder.istio_folder.id
}