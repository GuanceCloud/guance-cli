# Test import Grafana Dashboard into Guance Cloud

The details of CLI usage can be found in [reference documentations](../../../../docs/references/guance_iac_import_grafana.md).

![Preview](./grafana/node/preview.png)

## Import Dashboard

You can use Prometheus exporter and DataKit agent to collect and upload the metrics to Guance Cloud.

For example, the DataKit config is:

```toml
[[inputs.prom]]
  url = "http://node-exporter:9100/metrics"
  source = "prom"
  metric_types = []
  interval = "60s"
  measurement_name = "prom"
  [inputs.prom.tags]
    job = "DataKit"
```

Then download the [Node Exporter Dashboard on Grafana](https://grafana.com/grafana/dashboards/1860-node-exporter-full/).

Then run the Guance CLI to import the downloaded JSON.

* Run "guance iac import grafana -f ./iac/import/grafana/input.json -t terraform-module -o ./out"

You will get a Terraform module at `./out` folder. So you can apply it to create the real dashboard resources on Guance Cloud.

You can also see [./examples/node](config/node/) folder for a complete code example.
