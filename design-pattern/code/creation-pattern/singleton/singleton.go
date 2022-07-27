package singleton

import "sync"

type Singleton struct {
}

var instance *Singleton
var mutex = sync.Mutex{}
var once = sync.Once{}

// GetInstanceByDoubleCheck 利用mutex锁做双重检测
func GetInstanceByDoubleCheck() *Singleton {
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}
	return instance
}

// GetInstance 使用sync.Once.Do实现单例
func GetInstance() *Singleton {
	// sync.Once.Do可以保证在需要使用时，代码只执行一次
	// 内部使用done变量来记录，同时使用mutex和atomic来保证done变量的线程安全
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
