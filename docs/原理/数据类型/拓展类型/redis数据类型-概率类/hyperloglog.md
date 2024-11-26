# hyperloglog

## 算法原理

hyperloglog是一种**基数估算**算法

https://www.cnblogs.com/54chensongxia/p/13803465.html

https://blog.codinglabs.org/articles/algorithms-for-cardinality-estimation-part-i.html

## 主要用途

因为这个是一个大概估算，可以用在一些对精度不是特别准的场景，例如统计一个网站当天的访问量，100w和101w差别不大.

## 命令

内置命令，基于set