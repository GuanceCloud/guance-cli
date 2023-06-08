package prometheus

// Builder is the builder to convert Prometheus from Grafana Dashboard to Guance Cloud
type Builder struct {
	Measurement string
	ChartType   string
}
