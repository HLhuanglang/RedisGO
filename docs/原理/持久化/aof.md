#  aof持久化

鉴于rdb的耗时长且占用内存，所以引入aof. aof是指将写操作写入日志文件中，当redis重启时，把所有写操作执行一遍即可恢复数据。

这里引出了先写日志还是先写内存的问题，首先分析先写日志和先写内存优缺点：

- 先内存,后日志：速度快，但是更新日志时可能出现问题导致写操作丢失，进而导致数据丢失
- 先日志,后内存：数据/操作不会丢失，但是更新日志时需要先检测写命令是否正确

针对写操作，redis按照先写内存，然后写日志文件来进行持久化，这里我认为有几点原因：
- todo1
- todo2

## 如何实现AOF

redis写入aof日志的过程如下：
- 执行写操作命令后，将命令追加到server.aof_buf缓冲区
- **通过调用write系统调用，将aof_buf缓冲区的数据写入到aof文件中，注意此时写入的是内核缓冲区-->page_cache，需要等内核将数据写入磁盘**
- 具体内核缓冲区什么时候写入磁盘，有内核决定，但是内核提供了fsync、fdatasync等函数强制内核立即写回磁盘，针对这个写回时间，redis提供了三种策略
  - Always：同步写. 每次执行完写操作后，调用fsync同步写回磁盘
  - Everysec：每秒写. 每间隔1s调用fsync同步写回磁盘
  - No：os控制. 由操作系统控制什么时候写回磁盘

## 如何配置AOF

默认情况没有开启aof，需要手动修改redis.conf中配置项
```
#开启aof功能
appendonly on

#aof文件名称，默认是appendonly.aof
appendfilename "xxxx.aof"

#aof文件保存目录
dir ./

#写回策略
appendfsync always|everysec|no
```

## AOF重写