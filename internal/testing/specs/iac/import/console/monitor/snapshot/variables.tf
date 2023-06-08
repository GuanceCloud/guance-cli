variable "alert_policy_id" {
  type        = string
  default     = null
  description = "(Optional) The Alert Policy ID for taking effect when the alert is triggered."
}

variable "dashboard_id" {
  type        = string
  default     = null
  description = "(Optional) The Dashboard ID for linking to the report of monitoring."
}
