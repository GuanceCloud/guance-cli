# Guance Cloud CLI

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

### Quickstart

**Linux**

```shell
curl -sL https://raw.githubusercontent.com/GuanceCloud/guance-cli/master/install.sh | bash
```

See the [Release page](https://github.com/GuanceCloud/guance-cli/releases) to download the latest release for Linux.

**Mac OSX**

```shell
brew tap GuanceCloud/homebrew-tap
brew install GuanceCloud/tap/guance
```

**Windows**

We are working in progress.

**Verify Installation**

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

