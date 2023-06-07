C002: Grafana Importer
----
* Author(s): @yufeiminds
* Approver: @coanor
* Status: Draft
* Last updated: 2023-06-07

## Abstract

This proposal describes a component to import the Grafana dashboard into Guance Cloud.

The features are described as follows:

1. Build the chart models for Grafana and Guance Cloud.
2. Implement an extension mechanism to add new charts and fields converter.
3. Re-write the PromQL in Grafana to Guance Cloud PromQL/MetricsQL.
4. Add testing framework features to run integration testing for the Grafana dashboard.

The limitation:

1. Only support Prometheus Datasource. The graph with other datasource will return to an empty panel on Guance Cloud.
2. Not a full-featured conversion for Grafana. The supported charts and features are shown as follows.

The demo is like this:

![node exporter dashboard](./C002_images/node-exporter-dashboard.png)

## Proposal

### Dashboard model

Most of the work in this proposal is: **How to build the data model for Grafana and Guance Cloud**.
Because in software engineering, the static type is also a contract of system edge.
A determined type can make the Grafana importer's behavior exactly. 

1. The Guance Cloud model is described in the [json-model repository](https://github.com/GuanceCloud/json-model).
2. The Grafana model is described in the [grafana repository](https://github.com/grafana/grafana/blob/main/kinds/dashboard/dashboard_kind.cue).

### Feature list

All features of Grafana importer are listed here:

| Grafana Panel Type | Guance Chart Type | Implemented? | All Features Supported? |
| ------------------ |-------------------|--------------| ----------------------- |
| bar gauge          | bar               | YES          | [Partial](#BarGauge)    |
| gauge              | gauge             | YES          | [Partial](#Gauge)       |
| stat               | singlestat        | YES          | [Partial](#Stat)        |
| time series        | sequence          | YES          | [Partial](#TimeSeries)  |
| table              | table             | YES          | [Partial](#Table)       |
| pie                | pie               | NO           |                         |
| barchart           | bar               | NO           |                         |
| text               | text              | NO           |                         |
| candlestick        | -                 | -            |                         |
| flame graph        | -                 | -            |                         |
| geomap             | -                 | -            |                         |
| heatmap            | -                 | -            |                         |
| histogram          | -                 | -            |                         |
| logs panel         | -                 | -            |                         |
| state timeline     | -                 | -            |                         |
| status history     | -                 | -            |                         |
| traces             | -                 | -            |                         |

The `-` token means the panel couldn't convert to Guance Cloud directly,
such as it not based on prometheus data-source, don't have related chart type in Guance and so on.

#### BarGauge

Working in progress.

#### Gauge

Working in progress.

#### Stat

Working in progress.

#### TimeSeries

Working in progress.

#### Table

Working in progress.

### Re-write PromQL for Guance

According to [Guance Cloud PromQL] documentation (in chinese), Guance has a specific prefix about measurement on PromQL.

For example:

* The `node_network_receive_bytes_total` should be renamed as `node:network_receive_bytes_total`.

So Guance CLI will auto do this convert by PromQL AST when importing. You can also add customized measurement by CLI flag.

## Rationale

### Why not implement all the Grafana features at once?

Because it's too expensive to maintain.

We list features as much as possible. So make it as a long-term projects and has a public roadmap to push it.
Anyone who wants to use the particular chart could also contribute a new package to us or create an issue request for the maintainer team to add it.

## Implementation

1. Add schema management framework.
2. Implement Grafana Converter.
3. Implement the first batch charts, includes `gauge`, `stat`, `time series` and `table`.
4. Implement PromQL re-writer.
5. Integrate [Dashboard#1860](https://grafana.com/grafana/dashboards/1860-node-exporter-full/) with [DataKit](https://github.com/GuanceCloud/datakit), [Prometheus Node Exporter](https://github.com/prometheus/node_exporter) and Guance Cloud.
6. Improve integration testing framework, add dashboard integration.
