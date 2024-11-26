# Redis-lua脚本

参考：[Redis programmability | Docs](https://redis.io/docs/latest/develop/interact/programmability/)

## 命令

从 Redis 2.6.0 版本开始，通过内置的 Lua 解释器(redis内嵌的lua脚本版本是5.1)，可以使用 EVAL 命令对 Lua 脚本进行求值。在lua脚本中可以通过两个不同的函数调用redis命令，分别是：**redis.call()** 和 **redis.pcall()**

- redis.call()：在执行命令的过程中发生错误时，脚本会停止执行，并返回一个脚本错误，错误的输出信息会说明错误造成的原因.
- redis.pcall()： 出错时并不引发(raise)错误，而是返回一个带 err 域的 Lua 表(table)，用于表示错误



hast tag：在集群模式下，可以使用{}括起来，这样子能够保证所有的key都落在同一个机器上执行。



## redis-lua语法

需要注意的是，redis内置的lua版本是5.1，所以编写脚本时需要注意lua语法问题，最新的lua是5.4.

lua常见的数据类型：

![img](https://hl1998-1255562705.cos.ap-shanghai.myqcloud.com/Img/1031302-20201109145752957-801370564.png)

redis数据和lua数据转换：

![img](https://hl1998-1255562705.cos.ap-shanghai.myqcloud.com/Img/1031302-20201109145953858-1116129880.png)

![img](https://hl1998-1255562705.cos.ap-shanghai.myqcloud.com/Img/1031302-20201109150022613-1669952987.png)

## 调试脚本

### 使用LDB调试

从Redis 3.2开始，内置了 Lua debugger（简称LDB），使用Lua debugger可以很方便的对我们编写的Lua脚本进行调试。开启 lua dubegger ，将会进入debug命令行。
两种模式：
- 异步模式：`--ldb`。这个模式下 redis 会 fork 一个进程进入隔离环境，不会影响 redis 正常提供服务，但调试期间，原始 redis 执行命令、脚本的结果也不会体现到 fork 之后的隔离环境之中
- 同步模式：`--ldb-sync-mode`。同步模式，这个模式下，会阻塞 redis 上所有的命令、脚本，直到脚本退出，完全模拟了正式环境使用时候的情况，使用的时候务必注意这点

完整命令
```
redis-cli -h {redis_host} -p {redis_port} -a {redis_password} -n {database} --ldb eval  /path/to/script keys [key1 key2 key3…] , args [argv1 argv2 argv3…]
```


### 使用全局字符串

```lua
local g_debug_msg = ""
local function Debug(msg)
    g_debug_msg = g_debug_msg .. msg .. "\n"
end

-- 执行脚本,记录中间状态值
Debug("xxxxxx")

-- 将debug信息返回给调用
return g_debug_msg
```