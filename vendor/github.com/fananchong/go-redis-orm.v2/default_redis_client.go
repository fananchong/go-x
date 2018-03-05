package go_redis_orm

import (
	"github.com/fananchong/goredis"
)

type DefaultRedisClient struct {
	*goredis.Client
}

func NewDefaultRedisClient(dbName string, addrs []string, password string, dbindex int) (IClient, error) {
	this := &DefaultRedisClient{}
	option := goredis.NewDefaultOption()
	option.Password = password
	option.DBIndex = dbindex
	client, err := goredis.NewClient(dbName, addrs, option)
	if err != nil {
		return nil, err
	}
	this.Client = client
	return this, nil
}
