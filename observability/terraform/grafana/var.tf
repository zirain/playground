variable "gf_url" {
  description = "Grafana URL"
  type        = string
  default     = "http://localhost:3000"
}

variable "gf_auth" {
  description = "Grafana Auth"
  type        = string
  default     = ""
}

variable "threshold" {
  description = "Threshold for alerting"
  type = object({
    High4xxErrorRate = number
  })
  default = {
    High4xxErrorRate = 0.05
  }
}