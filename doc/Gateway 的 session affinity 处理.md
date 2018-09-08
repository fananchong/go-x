## 时序图

![Gateway 的 session affinity 处理](Gateway 的 session affinity 处理.svg)

## 说明

- gateway 失效时，可能导致的问题：

   2.2 步骤中 可能会把 新登陆请求的 SET NX 删除 （概率非常小）

   因此这里写的命令是 DELX 不是 Redis 原生的 DEL。需要用 redis module 做一个 `DELX` ：条件删除， if value = "xxx" then del key （或使用  Redis Lua 脚本）

- 6 步骤中的 EXPIRE 后，有小概率 可能会把 新登陆请求的 SET NX 删除

   因此步骤2.1 中写的是 SETX 不是 Redis 原生的 SET。需要用 redis module 做一个 `SETX` ： 不管 SET NX 有没有设置成功，都重置过期时间为 1 年 （或使用  Redis Lua 脚本）
