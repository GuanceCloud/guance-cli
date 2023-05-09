# Guance Cloud CLI

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat&logo=github&color=2370ff&labelColor=454545)](http://makeapullrequest.com)
[![Go Report Card](https://goreportcard.com/badge/github.com/GuanceCloud/guance-cli)](https://goreportcard.com/report/github.com/GuanceCloud/guance-cli)
[![Downloads](https://img.shields.io/github/downloads/GuanceCloud/guance-cli/total.svg)](https://github.com/GuanceCloud/guance-cli/releases)

A command-line tool to help users interact with Guance Cloud.

![cover](./artwork/cover.png)

## Features

1. Import external resource as Guance Cloud IaC resource (Terraform)
    * **Console**, see [these specs](specs/iac/import.spec.md) for more details.
    * **Grafana**, working in progress, is coming soon.
2. Continuation Integration/Testing Observability for GitHub
    * **JMeter**, working in progress, coming soon.
    * **JUnit**, working in progress, coming soon.

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

If you have any feedback, please create an issue or a pull request.

## Contributing

If you wish to contribute to this repository, please fork it and send us a pull request.

This [Contribution Guidelines](https://guance.io/contribution-guide/) document contains more detailed information about contributing to this repository.

## License

This repository is licensed under the [Apache 2.0 License](./LICENSE).
