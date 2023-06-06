FROM gcr.io/istio-testing/build-tools:master-62f8b507af660ecccb613e643daf5ea546222ac7

ENV GOLANGCI_LINT_VERSION=v1.53.2

RUN go install -ldflags="-s -w" github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}