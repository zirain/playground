
```shell
docker run -d -p 8080:80 kennethreitz/httpbin
```

```console
time curl -v -H "Content-Length: 15" localhost:9080/get

* Uses proxy env variable no_proxy == 'localhost,127.0.0.1,10.0.0.0/8,192.168.0.0/16,172.0.0.0/8'
*   Trying 127.0.0.1:9080...
* Connected to localhost (127.0.0.1) port 9080 (#0)
> GET /get HTTP/1.1
> Host: localhost:9080
> User-Agent: curl/7.81.0
> Accept: */*
> Content-Length: 15
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 408 Request Timeout
< content-length: 15
< content-type: text/plain
< date: Mon, 26 Feb 2024 08:28:13 GMT
< server: envoy
< connection: close
< 
* Closing connection 0
request timeout
real    0m0.112s
user    0m0.007s
sys     0m0.004s
```