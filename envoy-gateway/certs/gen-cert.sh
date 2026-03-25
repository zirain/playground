#!/bin/bash

set -euo pipefail

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=example Inc./CN=example.com' -keyout clientca.key -out clientca.crt
openssl req -new -newkey rsa:2048 -nodes -keyout client.key -out client.csr -subj "/CN=example-client/O=example organization"
openssl x509 -req -days 365 -CA clientca.crt -CAkey clientca.key -set_serial 0 -in client.csr -out client.crt
