copy form [wurstmeister/kafka-docker](https://github.com/wurstmeister/kafka-docker)

# pull images

```
docker pull wurstmeister/kafka
docker pull wurstmeister/zookeeper
docker pull kafkamanager/kafka-manager
```

# set up

```
docker-compose up -d
```

# kafka manager

```
http://localhost:9000/
```

# scale

```
docker-compose scale node=3
```

# shut down

```
docker-compose down
```