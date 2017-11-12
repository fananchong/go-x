### redis主从

  - install-redis.sh

  
### redis哨兵

  - install-redis-sentinel.sh
  
  

### Redis配置说明

  - timeout 0

  指定在一个 client 空闲多少秒之后关闭连接（0 就是不管它）


  - tcp-keepalive 60

  如果设置为非零，则在与客户端缺乏通讯的时候使用 SO_KEEPALIVE 发送 tcp acks 给客户端。


  - save ""

  存 DB 到磁盘。格式：save <间隔时间（秒）> <写入次数>


  - maxmemory 5gb

  最大使用内存
