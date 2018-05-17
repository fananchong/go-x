# go-discovery
discovery service for golang

### 功能

  - 封装etcd，提供事件触发接口
  - 节点自识别自身对外IP
  - 节点自动分配可用IP、多节点不会冲突


### 例子

```go
type MyNode struct {
	godiscovery.Node
}

func NewMyNode() *MyNode {
	this := &MyNode{}
	this.Node.Init(this)
	return this
}

func (this *MyNode) OnNodeUpdate(nodeIP string, nodeType int, id string, data []byte) {
	fmt.Println("OnNodeUpdate: nodeIP =", nodeIP, "nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *MyNode) OnNodeJoin(nodeIP string, nodeType int, id string, data []byte) {
	fmt.Println("OnNodeJoin: nodeIP =", nodeIP, "nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *MyNode) OnNodeLeave(nodeType int, id string) {
	fmt.Println("OnNodeLeave: nodeType =", nodeType, "id =", id)
}

func (this *MyNode) GetPutData() (string, error) {
	return string([]byte{1, 2, 3, 4}), nil
}

func main() {

	hosts := ""
	flag.StringVar(&hosts, "hosts", "101.132.47.70:12379,101.132.47.70:22379,101.132.47.70:32379", "etcd hosts")
	whatsmyip := ""
	flag.StringVar(&whatsmyip, "whatsmyip", "101.132.47.70:3000", "whatsmyip host")
	nodeType := 0
	flag.IntVar(&nodeType, "nodeType", 1, "node type")
	watchNodeTypes := ""
	flag.StringVar(&watchNodeTypes, "watchNodeTypes", "1,2,3,4", "watch node type")
	putInterval := int64(0)
	flag.Int64Var(&putInterval, "putInterval", 1, "put interval")

	flag.Parse()

	node := NewMyNode()
	node.OpenByStr(hosts, whatsmyip, nodeType, watchNodeTypes, putInterval)

	for {
		time.Sleep(time.Minute)
	}
}

```


### main函数中参数说明

  - hosts

    etcd 地址列表

  - whatsmyip

    whatsmyip 地址

  - nodeType

    自己节点的类型，非0，则定期(putInterval)调用GetPutData函数，上传数据

  - watchNodeTypes

    要监听的节点类型列表。非空，则OnNodeUpdate、OnNodeJoin、OnNodeLeave会被触发


  - putInterval

    调用GetPutData函数的频率，单位秒



### API使用注意事项

**OnNodeUpdate、OnNodeJoin、OnNodeLeave、GetPutData 在内部协程被调用，请注意多协程安全！！！**


### Etcd部署脚本说明

提供2种Docker Swarm方式部署etcd

启动脚本                                                                  | 说明
-------------------------------------------------------------------------|-----
docker-swarm/install-etcd-static.sh                                       | 静态配置方式部署etcd
docker-swarm/install-discovery.etcd.io.sh<br>docker-swarm/install-etcd.sh | etcd发现方式部署etcd


### WhatsMyIP部署脚本说明

启动脚本                                     | 说明
--------------------------------------------|-----
docker-swarm/install-whatsmyip.sh           | 启动whatsmyip服务
