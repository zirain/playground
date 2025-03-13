FROM envoyproxy/nighthawk-dev:latest AS source

FROM ubuntu:22.04@sha256:0e5e4a57c2499249aafc3b40fcd541e9a456aab7296681a3994d631587203f97

RUN apt-get update && apt-get install \
        autoconf \
        automake \
        cmake \
        curl \
        libtool \
        make \
        ninja-build \
        patch \
        python3-pip \
        unzip \
        virtualenv -y

COPY --from=source /usr/local/bin/nighthawk_client /usr/local/bin/nighthawk_client
COPY --from=source /usr/local/bin/nighthawk_test_server /usr/local/bin/nighthawk_test_server
COPY --from=source /usr/local/bin/nighthawk_output_transform /usr/local/bin/nighthawk_output_transform
COPY --from=source /usr/local/bin/nighthawk_service /usr/local/bin/nighthawk_service
COPY --from=source /usr/local/bin/nighthawk_adaptive_load_client /usr/local/bin/nighthawk_adaptive_load_client
COPY --from=source /etc/envoy/envoy.yaml /etc/envoy/envoy.yaml

# Ports for nighthawk_test_server, see default-config.yaml
EXPOSE 10001
EXPOSE 10080
# The default port for nighthawk_service
EXPOSE 8443

CMD ["-c", "/etc/envoy/envoy.yaml"]
