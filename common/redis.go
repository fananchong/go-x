package common

import (
	"github.com/fananchong/goredis"
)

type RedisObj struct {
	*goredis.Client
}

func NewRedisObj(dbname string, addrs []string) *RedisObj {
	this := &RedisObj{}
	option := goredis.NewDefaultOption()
	db, err := goredis.NewClient(dbname, addrs, option)
	if err != nil {
		GetLogger().Errorln(err)
		return nil
	}
	this.Client = db
	return this
}
