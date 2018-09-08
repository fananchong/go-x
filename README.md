# go-x

Go游戏服务器框架

## 编译

- Linux

  ```shell
  ./build.sh
  ```

- Windows

  ```shell
  build.bat
  ```

  _（build.bat中GOPATH是特殊路径，需要根据自己本地实际路径做修改）_


## 服务器运行

- Linux

  1. 安装 Kubernetes
  2. 执行 ./build.sh


- Windows

  1. 本地安装redis
  2. hosts文件增加 127.0.0.1 redis.go-x
  3. 执行 go-x\\example1_iogame\\bin\\start.bat


## H5Client 测试客户端

1. 更新代码

    ```shell
    git.exe submodule update --init -- "tools/h5client"
    ```

2. 在 go-x\\tools\\h5client 目录下执行 （只需一次）

    ```shell
    cnpm install
    ```

3. 在 go-x\\tools\\h5client 目录下执行

    ```shell
    run.bat
    ```

4. 登录界面服务器 IP 、 Port

  - Windows

    127.0.0.1 8080

  - Linux

    你服务器IP 30100



## 完成的功能

- Login 服务器

  - 帐号验证、登录


- Gateway 服务器

  - 消息中继


- Hub 服务器

  - 服务器组内消息广播


- Lobby 服务器

  - 创建角色信息
  - 查询角色信息


## 支持 kubernetes 部署

- Linux 完整支持 kubernetes 部署
- Windows 仅做单节点开发调试用

## TODO

- example1_iogame 从 go-x 中分离， go-x 将作为 example1_iogame 的 第三方库方式来使用

- Hub 重命名为 Center

- 实现 Gateway 的 session affinity 处理

- 实现 Lobby 的 session affinity 处理

- match匹配逻辑

- room实现一个简单的场景

- 其他

  - session发送缓冲区[]byte不必频繁创建（gotcp优化）
  - 增加服务器间相互ping操作
  - 待续
