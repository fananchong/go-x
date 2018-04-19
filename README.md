# go-x

Go游戏服务器框架

## 编译

```dos
git submodule update --init -- "tools/build"
build.bat
```

## 完成的功能

- Login服务器
- Gateway服务器
- Hub服务器

## TODO

- base使用例子

- match匹配逻辑

- room实现一个简单的场景

- 其他

  - 服务发现节点ID类型由string改为uint64
  - 负载均衡选取服务器时做过载(overload)判断
  - uid-gatewayId 键值对加过期时间，且gate且gateway定时刷新它
  - 待续
