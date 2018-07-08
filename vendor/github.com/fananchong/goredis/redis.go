package goredis

import (
	"errors"
	"strings"

	"github.com/gomodule/redigo/redis"
)

type IClient interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
}

type Client struct {
	cli IClient
}

func NewClient(dbName string, addrs []string, option *Option) (*Client, error) {
	this := &Client{}
	err := this.Init(dbName, addrs, option)
	return this, err
}

func (this *Client) Init(dbName string, addrs []string, option *Option) error {
	var err error = nil
	if len(addrs) <= 0 {
		return errors.New("addrs len error!")
	}
	switch option.Type {
	case Unknow:
		{
			var conn redis.Conn
			var isSentinel, isCluster string
			conn, err = getTempConn(addrs[0], option.Password)
			if err != nil {
				return err
			}
			defer conn.Close()
			isSentinel, err = redis.String(conn.Do("INFO", "Sentinel"))
			if err != nil {
				return err
			}
			isCluster, err = redis.String(conn.Do("INFO", "Cluster"))
			if err != nil {
				return err
			}

			if isSentinel != "" {
				this.cli = NewSentinelClient(dbName, addrs, option)
			} else {
				if strings.Contains(isCluster, "0") {
					this.cli = NewStandaloneClient(dbName, addrs[0], option)
				} else {
					this.cli, err = NewClusterClient(dbName, addrs, option)
				}
			}
		}
	case Standalone:
		{
			this.cli = NewStandaloneClient(dbName, addrs[0], option)
		}
	case Sentinel:
		{
			this.cli = NewSentinelClient(dbName, addrs, option)
		}
	case Cluster:
		{
			this.cli, err = NewClusterClient(dbName, addrs, option)
		}
	}
	return err
}

func (this *Client) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	return this.cli.Do(commandName, args...)
}

func getTempConn(addr string, password string) (redis.Conn, error) {
	c, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, err := c.Do("AUTH", password); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, nil
}
