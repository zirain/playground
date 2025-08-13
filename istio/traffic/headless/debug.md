# DNS proxying enabled

Got 503 when se is created without disabled IP

With DNS Proxying, pilot-agent(istio) capture the dns request and return `240.240.0.6`
the istio-proxy will try to connect `240.240.0.6` with authority `echo.sidecar` because 
`outbound|3000||echo.sidecar.svc.cluster.local` is ORIGNAL_DST type.

[2025-08-13T13:21:22.007Z] "GET / HTTP/1.1" 200 - via_upstream - "-" 0 413 1 1 "-" "curl/8.13.0" "471b668d-eae7-45a9-8446-27362051fc8c" "echo.sidecar:3000" "10.244.2.238:3000" outbound|3000||echo.sidecar.svc.cluster.local 10.244.3.173:39104 10.244.2.238:3000 10.244.3.173:39098 - default
[2025-08-13T13:25:07.651Z] "GET / HTTP/1.1" 503 URX,UF upstream_reset_before_response_started{connection_timeout} - "-" 0 114 30052 - "-" "curl/8.13.0" "7e812e50-0ecc-4047-9b3c-4dc104277597" "echo.sidecar:3000" "240.240.0.6:3000" outbound|3000||echo.sidecar.svc.cluster.local - 240.240.0.6:3000 10.244.3.173:52114 - default

# DNS proxying disabled

Request is fine with/without SE

[2025-08-13T13:27:46.394Z] "GET / HTTP/1.1" 200 - via_upstream - "-" 0 413 7 7 "-" "curl/8.13.0" "6e498106-9deb-4357-ab6a-07641e44decc" "echo.sidecar:3000" "10.244.3.176:3000" outbound|3000||echo.sidecar.svc.cluster.local 10.244.2.239:42594 10.244.3.176:3000 10.244.2.239:42578 - default
[2025-08-13T13:28:49.239Z] "GET / HTTP/1.1" 200 - via_upstream - "-" 0 413 15 14 "-" "curl/8.13.0" "bc9616f0-ae50-4a1c-9958-1dd97510632e" "echo.sidecar:3000" "10.244.3.176:3000" outbound|3000||echo.sidecar.svc.cluster.local 10.244.2.239:42594 10.244.3.176:3000 10.244.2.239:56150 - default