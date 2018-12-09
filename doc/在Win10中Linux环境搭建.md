## go-x 开发模式介绍

go-x 在 Windows10 中开发，并在 Linux 环境中调试运行

## go-x 脚本规范

go-x 下脚本将统一使用 Linux Shell 脚本

## Windows10 中 Linux 环境搭建过程

1. 安装 WSL

  - 控制面板->程序和功能->启用或关闭Windows功能->勾选 适用于Linux的Windows子系统

  - 打开应用商城搜索“WSL”，可根据自己需求选择安装一个或多个Linux系统
    - go-x 将使用安装 Ubuntu 18

  - Shift + 鼠标右键，点击`在此处打开 Linux Shell(L)` ,即可打开 WSL

2. 安装 Docker for Windows

  - 按照 https://www.docker.com/products/docker-desktop 中提示安装之

  - 打开 Docker for Windows - General ，勾选`Expose daemon on tcp://localhost:2375 without TLS`
  - 打开 Docker for Windows - Shared Drives ，勾选所有盘符
  - 打开 WSL ，安装 docker.io
    - 可以参考 https://blog.csdn.net/u013272009/article/details/81221661
  - 修正挂接目录问题
    - 可以参考 https://blog.csdn.net/u013272009/article/details/81222689


3. 安装 registry

  - 可以参考：https://github.com/fananchong/docker_script_memo/blob/master/registry.md


4. 安装 Kubernetes

  待续
