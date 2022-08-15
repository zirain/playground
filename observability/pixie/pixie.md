# Pixie

The Pixie platform consists of multiple components:

- **Pixie Edge Module (PEM)**: Pixie's agent, installed per node. PEMs use eBPF to collect data, which is stored locally on the node.
- **Vizier**: Pixie’s collector, installed per cluster. Responsible for query execution and managing PEMs.
- **Pixie Cloud**: Used for user management, authentication, and data proxying. Can be hosted or self-hosted.
- **Pixie CLI**: Used to deploy Pixie. Can also be used to run queries and manage resources like API keys.
- **Pixie Client API**: Used for programmatic access to Pixie (e.g. integrations, Slackbots, and custom user logic requiring Pixie data as an input)

![Pixie arch](./pixie.png)

## Pixie 101

### Network Monitoring

### Infrastructure Health

### Service Performance

### Database Query Profiling

### Request Tracing

### Continuous Application Profiling

### Kafka Monitoring


## Data Retention

***Pixie only retains up to 24 hours of data. To process longer time spans, plugins can integrate Pixie with long-term data solutions.***

## PxL Scripts

Pixie Language (PxL) is a domain-specific language for working with machine data, and uses a Python dialect. It is heavily influenced by the popular data processing library Pandas, and is almost a subset of Pandas. PxL is used by the Pixie Platform, allowing developers to create high performance data processing pipelines to monitor, secure and operate their applications and infrastructure.

Like Python, PxL is implicitly and strongly typed, supports high-level data types, and functions. Unlike Python, PxL is a dataflow language allows the Pixie platform to heavily optimize it's execution performance, while maintaining expressiveness for data processing. PxL programs are typically short-lived and have no implicit side effects. As a result, PxL has no support for classes, exceptions, other such features of Python

PxL can be executed by the Pixie platform by using either the web based UI, API or CLI.


## Integrations

- Slack
- Grafana
- OpenTelemetry


## Summary

|          | Pixie | tetragon |
| -------- | ----- | -------- |
| 安装难度 | 高    |          |
| 架构复杂 | 高    |          |
| 易用性   | 高    |          |
| 多集群   | 支持  |          |
| 可扩展性 | 高    |          |
