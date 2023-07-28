# Import Grafana Node Exporter Dashboard into Guance Cloud

![Preview](preview.png)

## Import by file

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
 
* Run "docker compose up -d"
* Run "guance iac import grafana -f ./input.json -t terraform-module -o ./out"

You will get a Terraform module at `./out` folder. So you can apply it to create the real dashboard resources on Guance Cloud.

---

* Run "docker compose down"
