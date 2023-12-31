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

	runTimes := 100000

	// Use the customized limited pool.
	p, _ := ants.NewPool(10)

	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Submit(syncCalculateSum)
	}
	// it will be blocked until all tasks are finished and submitted.
	fmt.Println("all tasks have been submitted.")
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks.\n")

}
