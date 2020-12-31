package main

import (
	"fmt"
	"search"
	"sync"
	"time"
)

func main() {
	MakeLoop()
	ParallelMakeLoop()
}

func ParallelMakeLoop() {
	start := time.Now().UnixNano()

	var wg sync.WaitGroup
	wg.Add(2)

	go func (wg *sync.WaitGroup) {
		for i := 0; i < 50; i++ {
			search.Make(i)
		}
		wg.Done()
	}(&wg)

	go func (wg *sync.WaitGroup) {
		for i := 50; i < 100; i++ {
			search.Make(i)
		}
		wg.Done()
	}(&wg)

	wg.Wait()
	fmt.Println("並列実行速度 秒", float64(time.Now().UnixNano() - start) / 1000000000)
}

func MakeLoop()  {
	start := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		search.Make(i)
		//search.Get()
	}
	//body := search.Get()
	//fmt.Println(body)
	fmt.Println("直列実行速度 秒", float64(time.Now().UnixNano() - start) / 1000000000)
}