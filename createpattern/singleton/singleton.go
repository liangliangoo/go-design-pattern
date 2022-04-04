package singleton

import "sync"

// Singleton 单例模式
type Singleton struct {

}

var singleton *Singleton
// sync.Once  精确一次
var once sync.Once

func GetInstance() *Singleton {
	// 保证里面的函数只能被执行一次
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
