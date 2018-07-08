package goredis

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)

type ClusterClient struct {
	cli    *redisc.Cluster
	dbName string
	option *Option
}

func NewClusterClient(dbName string, addrs []string, option *Option) (*ClusterClient, error) {
	cli := &ClusterClient{}
	err := cli.Init(dbName, addrs, option)
	return cli, err
}

func (this *ClusterClient) Init(dbName string, addrs []string, option *Option) error {
	this.dbName = dbName
	this.option = option
	this.cli = &redisc.Cluster{
		StartupNodes: addrs,
		DialOptions:  []redis.DialOption{redis.DialConnectTimeout(5 * time.Second)},
		CreatePool:   this.createPool,
	}
	if err := this.cli.Refresh(); err != nil {
		return err
	}
	return nil
}

func (this *ClusterClient) createPool(addr string, options ...redis.DialOption) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     this.option.PoolMaxIdle,
		MaxActive:   this.option.PoolMaxActive,
		Wait:        this.option.PoolWait,
		IdleTimeout: this.option.PoolIdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr, options...)
			if err != nil {
				return nil, err
			}
			if this.option.Password != "" {
				if _, err := c.Do("AUTH", this.option.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", this.option.DBIndex); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}

func (this *ClusterClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := this.cli.Get()
	if conn != nil {
		defer conn.Close()
		return conn.Do(commandName, args...)
	} else {
		return nil, errors.New("[redis] Can't get redis conn!")
	}
}
