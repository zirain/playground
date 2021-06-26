# mysql & docker

docker image: https://hub.docker.com/_/mysql

```shell

docker run --name cabinet-mysql -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:latest

```