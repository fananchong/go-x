# go-redis-orm.v2

本库通过定义json文件，使用工具生成redis orm类文件

可以处理1对1类型的数据、以及1对N类型的数据

## 1对1类型介绍

即1个ID对应一条数据。

工具生成的redis操作类，有以下API：

API         | 说明
------------|-----------------------------------
NewXXX      | 构建一个XXX redis操作类
Load        | 从数据库加载数据
Save        | 保存被改变的数据
Delete      | 删除本条数据
SetYYY      | 修改YYY属性
GetYYY      | 获取YYY属性
HasKey      | 是否有某条数据

1对1类型，redis操作类，映射到redis hashmap上。

每次保存，只更新被修改的字段。因此效率是很好的。


## 1对1类型数据例子

```go
func test11() {
	dbName := "db1"

	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)
	go_redis_orm.CreateDB(dbName, []string{"192.168.1.12:16379"}, "", 0)

	// key值为1的 TestStruct1 数据
	data1 := NewTestStruct1(dbName, 1)
	data1.SetMyb(true)
	data1.SetMyf1(1.5)
	data1.SetMyi5(100)
	data1.SetMys1("hello")
	data1.SetMys2([]byte("world"))
	err := data1.Save()
	if err != nil {
		panic(err)
	}

	data2 := NewTestStruct1(dbName, 1)
	err = data2.Load()

	if err == nil {
		if data2.GetMyb() != true ||
			data2.GetMyf1() != 1.5 ||
			data2.GetMyi5() != 100 ||
			data2.GetMys1() != "hello" ||
			string(data2.GetMys2()) != "world" {
			panic("#1")
		}
	} else {
		panic(err)
	}

	err = data2.Delete()
	if err != nil {
		panic(err)
	}

	var hasKey int
	hasKey, err = data2.HasKey()
	if hasKey != 0 {
		panic("#2")
	}

	fmt.Println("OK")
}
```


## 1对N类型介绍

即1个ID对应N条数据。

工具生成的redis操作类，除了上面的API，还有以下API：

API         | 说明
------------|-----------------------------------
NewItem     | 增加1条数据
GetItem     | 获取某条数据
DeleteItem  | 删除某条数据
GetItems    | 获取所有条数据

1对N类型，redis操作类，映射到redis hashmap上。

在hashmap上是这样的： key : itemId1-itemData1, itemId2-itemData2 ... itemIdN-itemDataN

考虑过把N条数据打包一起存储（即1对1的方式），也考虑过把每条数据全部打碎，变成hashmap的field。

这2种方法均存在性能上的问题。因此采用折中的方案：数据ID作为field，数据内容打包作为field的值。

这样每次保存，只更新被修改的那条数据。同时field数量相对的少1个数量级。

**友情提示：如果业务需求，N值很大，这样的数据是不适合使用redis存储的。最好把N限制在小于一两百条数据的应用场景。**


## 1对N类型的数据例子

```go
func test1n() {
	dbName := "db2"

	go_redis_orm.SetNewRedisHandler(go_redis_orm.NewDefaultRedisClient)
	go_redis_orm.CreateDB(dbName, []string{"192.168.1.12:16379"}, "", 0)

	data1 := NewTestStruct2(dbName, 8)
	item1 := data1.NewItem(1)
	item1.SetMyf2(99.9)
	item2 := data1.NewItem(2)
	item2.SetMys1("hello")
	item2.SetMys2([]byte("world"))
	err := data1.Save()
	if err != nil {
		panic(err)
	}

	data2 := NewTestStruct2(dbName, 8)
	err = data2.Load()
	if err != nil {
		panic(err)
	}
	fmt.Printf("2: %+v\n", data2.GetItem(1))
	fmt.Printf("2: %+v\n", data2.GetItem(2))
	data2.DeleteItem(1)
	data2.Save()

	data3 := NewTestStruct2(dbName, 8)
	data3.Load()
	for _, v := range data3.GetItems() {
		fmt.Printf("3: %+v\n", v)
	}
	data3.Delete()
	data3.Save()

	data4 := NewTestStruct2(dbName, 8)
	data4.Load()
	fmt.Printf("4: item count = %d\n", len(data4.GetItems()))

	fmt.Println("OK")
}
```


## 使用方法

  1. 定义json文件

     格式参考：example/redis_def/*.json

  1. 生成go文件，参考example/g.sh


## 编译

执行下列语句：

```shell
./build.sh
```

## SetNewRedisHandler接口

本库支持第3方redis客户端整合进本项目，通过调用go_redis_orm.SetNewRedisHandler函数

需要实现go_redis_orm.IClient接口
```go
type IClient interface {
	Do(commandName string, args ...interface{}) (reply interface{}, err error)
}
```

例子参考：[default_redis_client.go](default_redis_client.go)


## Redis单机、主从、哨兵、集群搭建

详细参见：http://blog.csdn.net/u013272009/article/details/78513251
