# SNI dynamic forward

```shell
curl https://www.qq.com -x localhost:10000 -k -v

wget -e use_proxy=yes -e http_proxy=http://localhost:10000 https://www.qq.com
```