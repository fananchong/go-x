package common

import (
	"github.com/fananchong/goredis"
)

type RedisObj struct {
	*goredis.Client
}

func NewRedisObj(dbname string, addrs []string) *RedisObj {
	GetLogger().Infoln("start connect redis, addrs =", addrs)
	this := &RedisObj{}
	option := goredis.NewDefaultOption()
	db, err := goredis.NewClient(dbname, addrs, option)
	if err != nil {
		GetLogger().Errorln(err)
		return nil
	}
	GetLogger().Infoln("connect redis success. addrs =", addrs)
	this.Client = db
	return this
}
