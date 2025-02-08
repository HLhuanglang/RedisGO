package lock

// DistributedLock 分布式锁接口
type DistributedLock interface {
	TryLock(key string) bool // 尝试获取锁，获取失败返回false
	Lock(key string)         // 获取锁，直到成功，会阻塞
	UnLock(key string) bool  // 释放锁
}
