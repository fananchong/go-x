# go-discovery
discovery service for golang


### 例子

```
type MyNode struct {
	godiscovery.Node
}

func NewMyNode() *MyNode {
	this := &MyNode{}
	this.Node.Init(this)
	return this
}

func (this *MyNode) OnNodeUpdate(nodeType int, id string, data []byte) {
	xlog.Infoln("OnNodeUpdate: nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *MyNode) OnNodeJoin(nodeType int, id string, data []byte) {
	xlog.Infoln("OnNodeJoin: nodeType =", nodeType, "id =", id, "data =", data)
}

func (this *MyNode) OnNodeLeave(nodeType int, id string) {
	xlog.Infoln("OnNodeLeave: nodeType =", nodeType, "id =", id)
}

func (this *MyNode) GetPutData() (string, error) {
	return "", nil
}

func main() {

	hosts := ""
	flag.StringVar(&hosts, "hosts", "192.168.1.4:12379,192.168.1.4:22379,192.168.1.4:32379", "etcd hosts")
	nodeType := 0
	flag.IntVar(&nodeType, "nodeType", 1, "node type")
	watchNodeTypes := ""
	flag.StringVar(&watchNodeTypes, "watchNodeTypes", "1,2,3,4", "watch node type")
	putInterval := int64(0)
	flag.Int64Var(&putInterval, "putInterval", 1, "put interval")

	flag.Parse()

	node := NewMyNode()
	node.SetLogger(xlog)
	node.OpenByStr(hosts, nodeType, watchNodeTypes, putInterval)

	for {
		time.Sleep(time.Minute)
	}
}
```


### API使用注意事项

**OnNodeUpdate、OnNodeJoin、OnNodeLeave、GetPutData 在内部协程被调用，请注意多协程安全！！！**


### Etcd部署脚本说明

  - 提供2种Docker Swarm方式部署etcd

启动脚本                                                                   | 说明
--------------------------------------------------------------------------|-----
docker-swarm/install-etcd-static.sh                                       | 静态配置方式部署etcd
docker-swarm/install-discovery.etcd.io.sh<br>docker-swarm/install-etcd.sh | etcd发现方式部署etcd


  - 封装etcd client api (v3)，使用例子请参见example目录
