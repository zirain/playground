# Install

```
docker-compose up -d
```

# Verify

```
curl -XGET https://localhost:9200 -u admin:admin --insecure
curl -XGET https://localhost:9200/_cat/nodes?v -u admin:admin --insecure
curl -XGET https://localhost:9200/_cat/plugins?v -u admin:admin --insecure
```

# Uninstall 

```
docker-compose down -v
```