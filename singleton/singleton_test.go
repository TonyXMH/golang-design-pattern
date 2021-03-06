package singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t*testing.T)  {
	ins1:=GetInstance()
	ins2:=GetInstance()
	if ins1 != ins2{
		t.Fatal("instance is not equal")
	}
}
const parCount = 100
func TestParallelSingleton(t *testing.T)  {
	start:=make(chan struct{})
	wg:=sync.WaitGroup{}
	wg.Add(100)
	instances:=[parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1]{
			t.Fatalf("instance is not equal")
		}
	}
}

