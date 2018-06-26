package godiscovery

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func (this *Node) NewNodeId() (uint32, error) {
	key := "__#ETCDID#__"
	rep, err := this.GetClient().Get(this.GetCtx(), key)
	if err != nil {
		return 0, err
	}

	var id uint32 = 0
	var version int64
	if rep.Count != 0 {
		temp, err := strconv.Atoi(string(rep.Kvs[0].Value))
		if err != nil {
			return 0, err
		}

		id = uint32(temp)
		version = rep.Kvs[0].Version
	} else {
		txnRep, err := this.GetClient().Txn(this.GetCtx()).
			If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
			Then(clientv3.OpPut(key, strconv.FormatInt(int64(id), 10))).
			Else(clientv3.OpGet(key)).
			Commit()
		if err != nil {
			return 0, err
		}
		if txnRep.Succeeded {
			version = 1
		} else {
			version = txnRep.Responses[0].GetResponseRange().Kvs[0].Version
		}
	}

	var data string
	for {
		id = id + 1
		if id == 0 {
			id = 1
		}
		data = strconv.FormatInt(int64(id), 10)
		txnRep, err := this.GetClient().Txn(this.GetCtx()).
			If(clientv3.Compare(clientv3.Version(key), "=", version)).
			Then(clientv3.OpPut(key, data)).
			Else(clientv3.OpGet(key)).
			Commit()
		if err != nil {
			return 0, err
		}
		if txnRep.Succeeded {
			break
		} else {
			version = txnRep.Responses[0].GetResponseRange().Kvs[0].Version
			id++
		}
		time.Sleep(time.Duration(rand.Int31n(5)+1) * time.Second)
	}
	return id, nil
}
