# Guance Cloud CLI

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-blue.svg?style=flat&logo=github&color=2370ff&labelColor=454545)](http://makeapullrequest.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/GuanceCloud/guance-cli)](https://goreportcard.com/report/github.com/GuanceCloud/guance-cli)
[![Downloads](https://img.shields.io/github/downloads/GuanceCloud/guance-cli/total.svg)](https://github.com/GuanceCloud/guance-cli/releases)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=GuanceCloud_guance-cli&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=GuanceCloud_guance-cli)

A command-line tool to help users interact with Guance Cloud.

![cover](./artwork/cover.png)

## Features

| Topic                                   | Feature           | Proposals                                                 | User Specification                       | Related Projects                                                                                                                                                                                                                                                                                                                                                                              |
| --------------------------------------- | ----------------- | --------------------------------------------------------- | ---------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 🔧 _Command-line Interface_             | **Core**          | [A001](./proposals/A001-guance-cli-overview.md)           | [View](specs/guance.spec.md)             | This repository                                                                                                                                                                                                                                                                                                                                                                               |
| 🚅 _Resource Exporter_                  | **Console**       | WIP                                                       | [View](specs/iac/import/console.spec.md) | [![terraform-guance-dashboard](https://img.shields.io/badge/guance-terraform--guance--dashboard-blue?style=flat-square&logo=github)](https://github.com/GuanceCloud/terraform-guance-dashboard)<br/>[![terraform-guance-monitor](https://img.shields.io/badge/guance-terraform--guance--monitor-blue?style=flat-square&logo=github)](https://github.com/GuanceCloud/terraform-guance-monitor) |
|                                         | **Grafana**       | WIP                                                       | [View](specs/iac/import/grafana.spec.md) | [![json-model](https://img.shields.io/badge/guance-json--model-blue?style=flat-square&logo=github)](https://github.com/GuanceCloud/json-model)                                                                                                                                                                                                                                                |
| 🚀 _Continuation Integration / Testing_ | **JMeter**        | [C001](./proposals/C001-continuous-integration-jmeter.md) | WIP                                      | [![jmeter](https://img.shields.io/badge/apache-jmeter-blue?style=flat-square&logo=github)](https://github.com/apache/jmeter)                                                                                                                                                                                                                                                                  |
|                                         | **JUnit**         | WIP                                                       | WIP                                      | [![junit](https://img.shields.io/badge/junit--team-junit5-blue?style=flat-square&logo=github)](https://github.com/junit-team/junit5)                                                                                                                                                                                                                                                          |
| 📦 _Components Installer_               | **DataKit**       | WIP                                                       | WIP                                      | [![DataKit](https://img.shields.io/badge/guance-DataKit-blue?style=flat-square&logo=github)](https://github.com/GuanceCloud/DataKit)                                                                                                                                                                                                                                                          |
|                                         | **SCheck**        | WIP                                                       | WIP                                      | [![SCheck](https://img.shields.io/badge/guance-SCheck-blue?style=flat-square&logo=github)](https://github.com/GuanceCloud/SCheck)                                                                                                                                                                                                                                                             |
| 🔭 _Ecosystem Integration_              | **GitHub Action** | -                                                         | WIP                                      | [![setup-guance](https://img.shields.io/badge/guance-setup--guance-blue?style=flat-square&logo=github)](https://github.com/GuanceCloud/setup-guance)                                                                                                                                                                                                                                          |
|                                         | **DevContainer**  | -                                                         | WIP                                      | WIP                                                                                                                                                                                                                                                                                                                                                                                           |

References:

1. For more details about the proposal governance mechanism, see [Guance CLI Proposals](./proposals/README.md).
2. For more details about user specification, see [Guance CLI User Specification](./specs/README.md).

## Installation

### Mac OSX

```shell
brew tap GuanceCloud/homebrew-tap
brew install GuanceCloud/tap/guance
```

### Ubuntu or Debian

```shell
echo "deb [trusted=yes] https://releases.guance.io/apt/ /" | sudo tee /etc/apt/sources.list.d/guance.list
sudo apt update
sudo apt install guance
```

### CentOS or RHEL

```shell
cat <<EOF | sudo tee /etc/yum.repos.d/guance.repo
[guance]
name=Guance Cloud Repo
baseurl=https://releases.guance.io/yum/
enabled=1
gpgcheck=0
EOF
sudo yum install -y guance
```

### Binary

See the [release page](https://github.com/GuanceCloud/guance-cli/releases) to download the latest release.

### Verify Installation

```shell
guance version
```

## Feedback

Please create an issue or pull request if you have any feedback.

## Contributing

If you wish to contribute to this repository, please fork it and send us a pull request.

This [Contribution Guidelines](https://guance.io/contribution-guide/) document contains more detailed information about contributing to this repository.

## License

This repository is licensed under the [Apache 2.0 License](./LICENSE).
