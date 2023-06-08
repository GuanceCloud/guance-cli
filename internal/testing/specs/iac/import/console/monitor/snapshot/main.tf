
resource "guance_monitor" "main_001" {
  manifest = file("${path.module}/manifest-001.json")

  alert_policy = var.alert_policy_id == null ? null : {
    id = var.alert_policy_id
  }

  dashboard = var.dashboard_id == null ? null : {
    id =  var.dashboard_id
  }
}

resource "guance_monitor" "main_002" {
  manifest = file("${path.module}/manifest-002.json")

  alert_policy = var.alert_policy_id == null ? null : {
    id = var.alert_policy_id
  }

  dashboard = var.dashboard_id == null ? null : {
    id =  var.dashboard_id
  }
}

