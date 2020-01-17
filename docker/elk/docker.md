# Pull image

`docker pull elasticsearch:7.0.1`

`docker pull kibana:7.0.1`

# Docker run with link

`docker run -d -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" --name elasticsearch elasticsearch:7.0.1`

`docker run -d -p 5601:5601 --link elasticsearch:es -e "ELASTICSEARCH_HOSTS=http://elasticsearch:9200" --name kibana kibana:7.0.1`

# Docker run with network

`docker network create es-net`

`docker network ls`

`docker run -d -p 9200:9200 -p 9300:9300 --net=es-net -e "discovery.type=single-node" --name elasticsearch elasticsearch:7.0.1`

`docker run -d -p 5601:5601 --net=es-net -e "ELASTICSEARCH_HOSTS=http://es:9200" --name kibana kibana:7.0.1`

`docker inspect es-net`

# Docker-compose

