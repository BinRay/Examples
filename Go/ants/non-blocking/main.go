package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

var sum int32

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	//fmt.Println("Hello World!")
}

func main() {
	defer ants.Release()

	pool, _ := NewTaskPool(10)

	runTimes := 1000000

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		pool.Submit(syncCalculateSum)
	}
	fmt.Println("all tasks have been submitted.")
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")

}
