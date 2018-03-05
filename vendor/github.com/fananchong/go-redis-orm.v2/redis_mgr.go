package go_redis_orm

import (
	"errors"
)

var g_redismgr *RedisMgr

type RedisMgr struct {
	dbs      map[string]IClient
	newRedis NewRedisType
}

func NewRedisMgr() *RedisMgr {
	return &RedisMgr{
		dbs: make(map[string]IClient),
	}
}

func (this *RedisMgr) Create(dbName string, addrs []string, password string, dbindex int) error {
	if _, ok := this.dbs[dbName]; ok {
		return nil
	}
	if this.newRedis == nil {
		return errors.New("no set new handler!")
	}
	db, err := this.newRedis(dbName, addrs, password, dbindex)
	if err != nil {
		return err
	}
	this.dbs[dbName] = db
	return nil
}

func (this *RedisMgr) Get(dbName string) IClient {
	if db, ok := this.dbs[dbName]; ok {
		return db
	}
	return nil
}

func (this *RedisMgr) SetNewRedisHandler(handler NewRedisType) {
	this.newRedis = handler
}

func init() {
	g_redismgr = NewRedisMgr()
}
