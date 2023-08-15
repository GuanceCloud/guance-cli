# Import Grafana Node Exporter Dashboard into Guance Cloud

![Preview](preview.png)

## Import by file

You can use Prometheus exporter and DataKit agent to collect and upload the metrics to Guance Cloud.

### Step 1: Configure the Prometheus

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

### Step 2: Download the Grafana Dashboard JSON

Then download the [Node Exporter Dashboard on Grafana](https://grafana.com/grafana/dashboards/1860-node-exporter-full/).

### Step 3: Run the DataKit to collect metrics from Prometheus

Then run the Guance CLI to import the downloaded JSON.

* Run "docker compose up -d"

### Step 4: Run the grafana importer to import the grafana dashboard

* Run "guance iac import grafana -f ./input.json -t terraform-module -o ./out"

You will get a Terraform module at `./out` folder.

### Step 5: Run the Terraform apply to create dashboard

* Run "cd ./out"
* Run "terraform init"
* Run "terraform apply"

So you can apply it to create the real dashboard resources on Guance Cloud.

Completed!

---

* Run "docker compose down"
