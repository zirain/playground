#!/bin/bash

GATEWAY_HOST=$(kubectl get gateway/example-gateway -o jsonpath='{.status.addresses[0].value}')
echo "send gRPC request to ${GATEWAY_HOST}:80"
grpcurl -plaintext -authority=www.example.com ${GATEWAY_HOST}:80 yages.Echo/Ping
echo "send gRPC request to ${GATEWAY_HOST}:443"
grpcurl --insecure -authority=www.example.com ${GATEWAY_HOST}:443 yages.Echo/Ping
echo "send gRPC request to ${GATEWAY_HOST}:4000"
grpcurl -insecure -authority=www.example.com ${GATEWAY_HOST}:4000 yages.Echo/Ping