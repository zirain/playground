resource "grafana_rule_group" "dataplane" {
  name             = "Istio High 4xx Error Rate"
  folder_uid       = grafana_folder.istio_folder.uid
  interval_seconds = 60
  org_id           = 1

  rule {
    name           = "Istio High 4xx Error Rate"
    condition      = "C"
    for            = "1m"
    no_data_state  = "OK"
    exec_err_state = "Error"


    // Query the datasource.
    data {
      ref_id = "A"
      relative_time_range {
        from = 600
        to   = 0
      }
      datasource_uid = grafana_data_source.prometheus.uid
      model = jsonencode({
        editorMode = "builder"
        expr       = "sum by (destination_app, source_app, instance) (rate(istio_requests_total{response_code=~\"4.*\"}[5m]))/sum by (destination_app, source_app, instance) (rate(istio_requests_total[5m])) > ${var.threshold.High4xxErrorRate}"
        refId      = "A"
      })
    }

    data {
      datasource_uid = "__expr__"
      model = jsonencode({
        conditions = [
          {
            evaluator = {
              params = []
              type   = "gt"
            }
            operator = {
              type = "and"
            }
            query = {
              params = ["B"]
            }
            reducer = {
              params = []
              type   = "last"
            }
            type = "query"
          }
        ]
        expression = "A"
      })
      ref_id = "B"
      relative_time_range {
        from = 300
        to   = 0
      }
    }

    data {
      datasource_uid = "__expr__"
      ref_id         = "C"
      relative_time_range {
        from = 0
        to   = 0
      }
      model = jsonencode({
        conditions = [
          {
            evaluator = {
              params = [0]
              type   = "gt"
            }
            operator = {
              type = "and"
            }
            query = {
              params = ["C"]
            }
            reducer = {
              params = []
              type   = "last"
            }
            type = "query"
          }
        ]
      })
    }
  }
}
