https://cloud.tencent.com/document/product/239/89468

## 大key


### 是什么

对于不同的数据类型，界定存在差异，通常以key的大小和key中成员的数量来判断：
- string：value值超过10MB
- set：成员数量太多 >=10000
- zset：成员数量太多 >=10000
- list：成员数量太多 >=10000
- hash：成员数量太多 >=10000

### 有什么问题

- 大key会导致redis使用内存飙升，如果触发maxmemory策略，可能会导致重要key被移除
- 请求响应时间增加（操作大key耗时，可能阻塞）
- 同步中断or主从切换：例如对大key进行rename，容易造成主库长时间阻塞，进而可能引发同步中断或者主从切换
- 网络拥塞：如果频繁读取大key，会导致带宽被占满

### 为什么会出现
- 使用string类型时，存放大体积的二进制数据
- 业务设计不合理，导致某些key的成员过多
- 未定时清理数据，例如hash类型中key的成员越来越多
- 使用list类型，只增不减

### 如何定位
- 使用redis-cli命令：bigkeys、memkeys
    - 优点：方便、快速、安全
    - 缺点：分析结果不可定制化，时效性差；需要遍历所有key，可能影响性能
- 通过内置命令对可以进行分析：例如string使用strlen查询长度
    - 优先：对线上服务影响小
    - 缺点：不够准确，仅供参考
- 使用redis-rdb-tools工具

### 如何解决
- 对value进行压缩
- 进行拆分
- 数据清理，考虑其他方式来存储
    - redis 4.0之后，使用unlink命令可以非阻塞删除大key
    - redis 4.0之前，建议先使用scan扫描，然后进行删除，避免一次性删除大量key导致redis阻塞
- 过期数据清理

### 如何预防
- 配置实例告警

## 热key

### 是什么

### 有什么问题

### 如何解决