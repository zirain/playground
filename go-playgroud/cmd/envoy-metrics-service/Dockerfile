FROM ubuntu:latest
WORKDIR /root/

COPY envoy-metrics-service .
RUN chmod +x envoy-metrics-service

CMD ["./envoy-metrics-service"]