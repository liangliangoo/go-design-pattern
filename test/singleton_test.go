package test

import (
	"go-design-pattern/createpattern/singleton"
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := singleton.GetInstance()
	ins2 := singleton.GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	// 类似于java中的CountDownLatch
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*singleton.Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = singleton.GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
