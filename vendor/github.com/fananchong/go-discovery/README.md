# go-discovery
discovery service for golang


### 功能

  - 提供2种Docker Swarm方式部署etcd

启动脚本                                                                   | 说明
--------------------------------------------------------------------------|-----
docker-swarm/install-etcd-static.sh                                       | 静态配置方式部署etcd
docker-swarm/install-discovery.etcd.io.sh<br>docker-swarm/install-etcd.sh | etcd发现方式部署etcd


  - 封装etcd client api (v3)，使用例子请参见example目录
