# 配置文件

## 1.配置详解

### 1.1 INCLUDE

include /path/to/local.conf

可以导入其他配置文件，注意config rewrite不会修改include选项

### 1.2 MODULE

加载自定义模块(一般实现一些命令字、数据结构等)

loadmodule /path/to/my_module.so

### 1.3 NETWORK

需要注意一下tcp-backlog。tcp连接有syns queue和accept queue
syns queue：/proc/sys/net/ipv4/tcp_max_syn_backlog
accept queue：min(backlog, /proc/sys/net/core/somaxconn)，其中backlog是listen时候传入的，这个值的设置有点讲究
可以使用ss -lt看Send-Q来确认下当前accept queue的大小

参考：https://www.cnxct.com/something-about-phpfpm-s-backlog/

### 1.4 GENERAL

### 1.5 SNAPSHOTTING

### 1.6 REPLICATION

### 1.7 SECURITY

### 1.8 CLIENTS

### 1.9 MEMORY MANAGEMENT

内存管理

## 2.命令

Redis使用config命令，可以对配置项参数热修改，不必重启，重启可能引发一些问题：
- 数据丢失
    - 无持久化的实例：虽然默认是RDB方式
    - 延迟同步的数据：在主从架构中，如果主节点在重启的时候未同步到从节点，可能会导致这些数据在从节点上丢失
- 连接中断
    - 客户端与redis之间的连接会中断，要确保客户端是否又重连的逻辑。
- 性能影响
    - 重启时间：海量数据情况下，重新读rdb和aof耗时
    - 资源消耗：cpu和内存资源高峰(如果其他服务同机部署，会有影响)
- 配置不一致
    - 如果使用config set命令动态修改了配置而未使用config rewrite保存更改，那么重启后读到旧配置，可能导致实际行为与预期不符合
- 高可用性问题
    - TODO，与搭建的redis服务架构有关

具体命令字：
- config get：获取当前运行使用的配置
- config set：
    - 动态修改(修改内存中的内容,而非redis.conf)，无需重启即可生效
- config rewrite：
    - 它会生成一个新的配置文件，反映当前的设置，重写保存到redis.conf
    - 会持久化当前的配置，使其在 Redis 重启后依然有效。重写过程会尽量保留原始配置文件的结构和注释。