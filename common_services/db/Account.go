/// -------------------------------------------------------------------------------
/// THIS FILE IS ORIGINALLY GENERATED BY redis2go.exe.
/// PLEASE DO NOT MODIFY THIS FILE.
/// -------------------------------------------------------------------------------

package db

import (
	"errors"

	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/gomodule/redigo/redis"
)

type Account struct {
	Key  string
	pswd string
	uid  uint64

	__dirtyData map[string]interface{}
	__isLoad    bool
	__dbKey     string
	__dbName    string
	__expire    uint
}

func NewAccount(dbName string, key string) *Account {
	return &Account{
		Key:         key,
		__dbName:    dbName,
		__dbKey:     "Account:" + key,
		__dirtyData: make(map[string]interface{}),
	}
}

// 若访问数据库失败返回-1；若 key 存在返回 1 ，否则返回 0 。
func (this *Account) HasKey() (int, error) {
	db := go_redis_orm.GetDB(this.__dbName)
	val, err := redis.Int(db.Do("EXISTS", this.__dbKey))
	if err != nil {
		return -1, err
	}
	return val, nil
}

func (this *Account) Load() error {
	if this.__isLoad == true {
		return errors.New("alreay load!")
	}
	db := go_redis_orm.GetDB(this.__dbName)
	val, err := redis.Values(db.Do("HGETALL", this.__dbKey))
	if err != nil {
		return err
	}
	if len(val) == 0 {
		return go_redis_orm.ERR_ISNOT_EXIST_KEY
	}
	var data struct {
		Pswd string `redis:"pswd"`
		Uid  uint64 `redis:"uid"`
	}
	if err := redis.ScanStruct(val, &data); err != nil {
		return err
	}
	this.pswd = data.Pswd
	this.uid = data.Uid
	this.__isLoad = true
	return nil
}

func (this *Account) Save() error {
	if len(this.__dirtyData) == 0 {
		return nil
	}
	db := go_redis_orm.GetDB(this.__dbName)
	if _, err := db.Do("HMSET", redis.Args{}.Add(this.__dbKey).AddFlat(this.__dirtyData)...); err != nil {
		return err
	}
	if this.__expire != 0 {
		if _, err := db.Do("EXPIRE", this.__dbKey, this.__expire); err != nil {
			return err
		}
	}
	this.__dirtyData = make(map[string]interface{})
	return nil
}

func (this *Account) Delete() error {
	db := go_redis_orm.GetDB(this.__dbName)
	_, err := db.Do("DEL", this.__dbKey)
	if err == nil {
		this.__isLoad = false
		this.__dirtyData = make(map[string]interface{})
	}
	return err
}

func (this *Account) IsLoad() bool {
	return this.__isLoad
}

func (this *Account) Expire(v uint) {
	this.__expire = v
}

func (this *Account) GetPswd() string {
	return this.pswd
}

func (this *Account) GetUid() uint64 {
	return this.uid
}

func (this *Account) SetPswd(value string) {
	this.pswd = value
	this.__dirtyData["pswd"] = value
}

func (this *Account) SetUid(value uint64) {
	this.uid = value
	this.__dirtyData["uid"] = value
}
