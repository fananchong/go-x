package goredis

import "time"

type RedisType int

const (
	Unknow RedisType = iota
	Standalone
	Sentinel
	Cluster
)

type Option struct {
	Type            RedisType
	Password        string
	DBIndex         int
	PoolMaxIdle     int
	PoolMaxActive   int
	PoolWait        bool
	PoolIdleTimeout time.Duration
}

func NewDefaultOption() *Option {
	return &Option{
		Type:            Unknow,
		Password:        "",
		DBIndex:         0,
		PoolMaxIdle:     10,
		PoolMaxActive:   0,
		PoolWait:        true,
		PoolIdleTimeout: 240 * time.Second,
	}
}
