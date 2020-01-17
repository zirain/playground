# Pull image

`docker pull cosnul`

# 安装单个Consul

````
docker run -d -p 8500:8500 -e CONSUL_BIND_INTERFACE='eth0' --name consul_server_1 consul
````

## 查看当前集群信息

````
docker exec consul_server_1 consul members
````

## 新增Consul Server节点加入集群

````
docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_2 consul agent -server -join='172.17.0.2'

docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_server_3 consul agent -server -join='172.17.0.2'
````

## 新增Consul Client节点加入集群

为了减少网络损耗，每个应用服务部署Client节点

````
docker run -d -e CONSUL_BIND_INTERFACE='eth0' --name=consul_client_1 consul agent -join='172.17.0.2' -client='0.0.0.0'
````