C001: Continuous Integration JMeter
----
* Author(s): @yufeiminds
* Approver: @coanor
* Status: In Review
* Last updated: 2023-05-14
* Discussion at: GitHub issues (comma separated, filled after thread exists)

## Abstract

This proposal describes a component to upload JMeter Logs to Guance Cloud.

![JMeter Demo](./C001_images/jmeter-dashboard.png)

It will be used in continuous integration pipelines, such as GitHub Action, GitLab CI, etc.

## Background

JMeter is a popular tool for load testing. It is widely used in the continuous integration pipelines.

In real-world usages, there are two use cases for JMeter:

1. **Load generation**. In this use case, JMeter is a long-time running process. It will send pre-defined requests to the target system. And then record the response time and other metrics. The metrics will be used to evaluate the performance of the target system.
2. **Acceptance testing**. In this use case, JMeter is invoked by the continuous integration pipeline. It is mainly used to verify the correctness of the target system and compare the JMeter metrics changes between the source code version from version control systems such as Git.

## Proposal

DataKit Plugin will implement the **load generation** use case. Because DataKit is a long-time running process. It is more suitable to collect JMeter logs living.

This proposal focuses on the **acceptance testing** use case.

### Command-Line Arguments

We should implement a new sub-command `guance ci upload jmeter` to upload JMeter logs to Guance Cloud.

|Argument|Description|
|:--|:--|
| `tags` | The tags of the request. <br/> **Examples**: `--tags k1:v1 --tags k2:v2` |
| `service` | The service name. <br/> **Examples**: `foo` |
| `env` | The environment identifier. <br/> **Examples**: `production` |

### Collect JMeter logging

The JMeter logs contain the following tags:

|Tag|Description|
|:--|:--|
| `label` | The label of the request. <br/> **Examples**: `Upload Metrics Data`|
| `responseCode` | The response code of the request. <br/> **Examples**: `200`|
| `responseMessage` | The response message of the request. <br/> **Examples**: `OK`|
| `threadName` | The thread name of the request. <br/> **Examples**: `Thread Group 1-1`|
| `dataType` | The data type of the request. <br/> **Examples**: `text`|
| `success` | The success of the request. <br/> **Examples**: `true`|
| `failureMessage` | The failure message of the request. <br/> **Examples**: `Test failed: code expected to contain /200/`|

The JMeter logs contain the following fields:

|Field|Description|
|:--|:--|
| `elasped` | The elapsed time of the request. <br/> **Examples**: `2615`|
| `bytes` | The bytes of the request. <br/> **Examples**: `496`|
| `sentBytes` | The sent bytes of the request. <br/> **Examples**: `608`|
| `grpThreads` | The group threads of the request. <br/> **Examples**: `147`|
| `allThreads` | All threads of the request. <br/> **Examples**: `147` |
| `URL` | The URL of the request. <br/> **Examples**: `https://openway.guance.com/v1/write/metric?token=...`|
| `Latency` | The latency of the request. <br/> **Examples**: `2610`|
| `IdleTime` | The idle time of the request. <br/> **Examples**: `0`|
| `Connect` | The connect time of the request. <br/> **Examples**: `1665`|

The typical JMeter logging example is shown below:

```csv
timeStamp,elapsed,label,responseCode,responseMessage,threadName,dataType,success,failureMessage,bytes,sentBytes,grpThreads,allThreads,URL,Latency,IdleTime,Connect
1662531305870,2615,Upload Metrics Data,200,OK,Testing 1-1,,true,,496,608,147,147,https://openway.guance.com/v1/write/metric?token=...,2610,0,1665
1667975613816,408,Login,Non HTTP response code: java.net.UnknownHostException,Non HTTP response message: testing-ft2x-auth-api.cloudcare.cn: Name does not resolve,Testing 1-2,text,false,,2341,0,1,1,http://testing-ft2x-auth-api.cloudcare.cn/api/v1/auth/signin,0,0,408
```

It is a CSV format. The first line is the header. The following lines are the JMeter logging.

### Collect Git-related information

The Git-related information contains the following tags:

|Tag|Description|
|:--|:--|
|`git_repository_url`|The URL of the git repository.<br/> **Examples**: `https://github.com/GuanceCloud/community.git`|
|`git_branch`|The branch of the git repository.<br/> **Examples**: `main`|
|`git_tag`|The tag of the git repository.<br/> **Examples**: `v1.0.0`|
|`git_commit_sha`|The commit sha consists of the git repository.<br/> **Examples**: `b1b9c2c`|
|`git_message`|The message of the git repository.<br/> **Examples**: `docs: Add JMeter documentation`|
|`git_author_name`|The author name of the git repository.<br/> **Examples**: `guance-bot`|
|`git_author_email`|The author email of the git repository.<br/> **Examples**: `developer@guance.com` |
|`git_author_date`|The author date of the git repository.<br/> **Examples**: `2021-05-14T08:00:00Z`|
|`git_committer_name`|The committer name of the git repository.<br/> **Examples**: `guance-bot`|
|`git_committer_email`|The committer email of the git repository.<br/> **Examples**: `developer@guance.com` |
|`git_committer_date`|The committer date of the git repository.<br/> **Examples**: `2021-05-14T08:00:00Z`|

### Collect Runtime-related information

The Runtime-related information contains the following tags:

|Tag|Description|
|:--|:--|
|`platform`|The platform of the environment.<br/> **Examples**: `windows`, `darwin`, `linux`|
|`arch`|The architecture of the environment.<br/> **Examples**: `amd64`, `arm64`|
|`os`|The operating system of the environment.<br/> **Examples**: `windows 10`, `macos 11.3.1`, `ubuntu 20.04`|

## Rationale

### What are the differences between this Guance CLI and DataKit JMeter Plugin?

The running time is the crucial difference between the Guance CLI and DataKit Plugin.

* The Guance CLI is a short-time running process. It will only execute once in the continuous integration pipelines. And you can easily integrate it with most of the open-source CI pipeline ecosystems, such as GitHub Action, GitLab CI, etc.
* The DataKit Plugin is a long-time running process. In most use cases, it will collect the metrics by the scheduler or stream living. It will keep running in any client workload, such as Virtual Machine, Kubernetes, etc.

So in a real-world use case. We generally use the Guance CLI and DataKit JMeter Plugin both. The CLI is used in the development stage, helping us to manage the **SDLC (Software Development Life Cycle)**. The DataKit Plugin is used in the production stage to help us to manage the **SLO (Service Level Objective)**. We use both in Guance Cloud to help us manage all the metrics data and archive **Testing Observability**.

## Implementation

1. Create a new component in the Guance CLI repository.
2. Create a GitHub Action to set up the Guance CLI.
3. Create an interactive example to show how to use the component.
