package db

import (
	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/gomodule/redigo/redis"
)

const (
	SUID_TYPE_ACCOUNT = "account"
)

// Server Unique Identifier
type SUID struct {
	Cli go_redis_orm.IClient
}

func (this *SUID) New(typ string) (uint64, error) {
	return redis.Uint64(this.Cli.Do("HINCRBY", "suid", typ, 1))
}
