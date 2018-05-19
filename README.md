# go-x

Go游戏服务器框架

## 编译

```dos
git submodule update --init -- "tools/build"
build.bat
```

_（build.bat中GOPATH是特殊路径，所以你下载下来不一定编译通过）_

## 完成的功能

- Login服务器
- Gateway服务器
- Hub服务器

## TODO

- lobby使用例子

  - 客户端查询玩家基本信息，返回玩家名字

- match匹配逻辑

- room实现一个简单的场景

- 其他

  - Gateway转发客户端消息时，附带上UID信息
  - 负载均衡选取服务器时做过载(overload)判断
  - uid-gatewayId 键值对加过期时间，且gate且gateway定时刷新它
  - session发送缓冲区[]byte不必频繁创建（gotcp优化）
  - 待续
