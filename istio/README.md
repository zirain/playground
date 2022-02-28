# local development

https://github.com/howardjohn/local-istio-development

# istio devstats

https://istio.teststats.cncf.io/


# PR dashboard

https://k8s-gubernator.appspot.com/pr


# How to test istio.io repo using Kind

```
TEST_ENV=kind ADDITIONAL_CONTAINER_OPTIONS="--network host"  make doc.test TEST=tasks/observability/logs/otel-provider
```