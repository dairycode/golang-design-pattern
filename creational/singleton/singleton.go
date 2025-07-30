package singleton

import "sync"

type Singleton struct{}

// 饿汉式单例
var eager *Singleton

func init() {
	eager = &Singleton{}
}

func GetEager() *Singleton {
	return eager
}

// 懒汉式单例
var (
	lazy *Singleton
	once sync.Once
)

func GetLazy() *Singleton {
	once.Do(func() {
		lazy = &Singleton{}
	})
	return lazy
}
