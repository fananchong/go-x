# go-x

Go游戏服务器框架

## 编译

- linux

  ```shell
  build.sh
  ```

- windows

  ```shell
  build.bat
  ```

  _（build.bat中GOPATH是特殊路径，需要根据自己本地实际路径做修改）_

## 完成的功能

- Login服务器
- Gateway服务器
- Hub服务器

## 支持 kubernetes 部署

- linux 直接支持 kubernetes 部署
- windows 仅做单节点开发调试用

## TODO

- lobby使用例子

  - 客户端查询玩家基本信息，返回玩家名字

  
- match匹配逻辑

- room实现一个简单的场景

- 其他

  - Gateway转发客户端消息时，附带上UID信息
  - uid-gatewayId 键值对加过期时间，且gate且gateway定时刷新它
  - session发送缓冲区[]byte不必频繁创建（gotcp优化）
  - 增加服务器间相互ping操作
  - 待续
