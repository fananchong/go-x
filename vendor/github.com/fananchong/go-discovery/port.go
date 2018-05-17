package godiscovery

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net"
	"strconv"

	"github.com/coreos/etcd/clientv3"
)

type Port struct {
	port      uint16
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func (this *Port) GetPort() uint16 {
	return this.port
}

func (this *Port) Init(root context.Context, client *clientv3.Client) error {
	key := "__#etcdport#__"
	this.ctx, this.ctxCancel = context.WithCancel(root)
	rep, err := client.Get(this.ctx, key)
	if err != nil {
		return err
	}

	var port uint16 = 1024
	var pre uint = 0
	if rep.Count != 0 {
		temp, err := strconv.Atoi(string(rep.Kvs[0].Value))
		if err != nil {
			return err
		}

		port = uint16(uint(temp) % math.MaxUint16)
		pre = uint(temp) - uint(port)
	} else {
		_, err := client.Txn(this.ctx).
			If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
			Then(clientv3.OpPut(key, strconv.Itoa(int(port)))).
			Commit()
		if err != nil {
			return err
		}
	}

	for {
		port = this.getVaildPort(port)
		if port == 0 {
			return errors.New("invild port!")
		}
		data := strconv.Itoa(int(pre) + int(port))
		txnRep, err := client.Txn(this.ctx).
			If(clientv3.Compare(clientv3.Value(key), "<", data)).
			Then(clientv3.OpPut(key, data)).
			Commit()
		if err != nil {
			return err
		}
		if txnRep.Succeeded {
			break
		} else {
			port++
		}
	}
	this.port = port
	fmt.Printf("node's port:%d\n", port)
	return nil
}

func (this *Port) getVaildPort(port uint16) uint16 {
	counter := 0
	for {
		counter++
		if counter > math.MaxUint16+100 {
			break
		}

		port = port + 1
		if port == 0 {
			continue
		}
		address := fmt.Sprintf(":%d", port)
		tcpAddr, err := net.ResolveTCPAddr("tcp", address)
		if err != nil {
			continue
		}
		listener, err := net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			if listener != nil {
				listener.Close()
			}
			continue
		}
		listener.Close()
		return port
	}
	return 0
}
