output "monitor_ids" {
  description = "The `guance_monitor`'s id."
  value = [
    guance_monitor.main_001.id,
    guance_monitor.main_002.id,
  ]
}
