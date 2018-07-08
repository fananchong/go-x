package goredis

import (
	"errors"
	"time"

	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
)

type SentinelClient struct {
	masters    *redis.Pool
	masterName string
}

func NewSentinelClient(masterName string, addrs []string, option *Option) *SentinelClient {
	cli := &SentinelClient{}
	cli.Init(masterName, addrs, option)
	return cli
}

func (this *SentinelClient) Init(masterName string, addrs []string, option *Option) {
	sntnl := &sentinel.Sentinel{
		Addrs:      addrs,
		MasterName: masterName,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	this.masterName = masterName
	this.masters = &redis.Pool{
		MaxIdle:     option.PoolMaxIdle,
		MaxActive:   option.PoolMaxActive,
		Wait:        option.PoolWait,
		IdleTimeout: option.PoolIdleTimeout,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			if option.Password != "" {
				if _, err := c.Do("AUTH", option.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", option.DBIndex); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if !sentinel.TestRole(c, "master") {
				return errors.New("[redis] Role check failed")
			} else {
				return nil
			}
		},
	}
}

func (this *SentinelClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := this.masters.Get()
	if conn != nil {
		defer conn.Close()
		return conn.Do(commandName, args...)
	} else {
		return nil, errors.New("[redis] Can't get redis conn!")
	}
}
