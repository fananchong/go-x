# h5client

H5 测试客户端

## 环境搭建

安装以下工具：

- nodejs

  点击官方页面下载安装：[nodejs官网](https://nodejs.org)

- cnpm

  ```dos
  npm install cnpm -g --registry=https://registry.npm.taobao.org
  ```

- bower

  ```dos
  cnpm install -g bower
  ```

  或者

  ```dos
  npm install -g bower
  ```

- gulp

  ```dos
  cnpm install -g gulp
  ```

  或者

  ```dos
  npm install -g gulp
  ```

## 更新依赖

```
  cnpm install
  bower install
```

## 编译、打包

- 网页开发版

  ```dos
  gulp dev
  ```

  编译到build目录下，点击index.html来访问。

- 网页发布版

  ```dos
  gulp build
  ```

  编译到build目录下，点击index.html来访问。

- NodeJS服务器版

  ```
  gulp serve
  ```

  浏览器键入`http://127.0.0.1:8000`，来访问。

- 桌面版

  ```
  gulp dist
  ```

  发布在.temp-dist目录下。
