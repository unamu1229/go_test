package main

import (
	"fmt"
	"os"
	"search"
	"sync"
	"time"
)

func main(){
	GetLoop()
	if len(os.Args) > 1 {
		if os.Args[1] == "get" {
			GetLoop()
		}

		if os.Args[1] == "put" {
			MakeLoop()
			ParallelMakeLoop()
		}
	}
}

func GetLoop() {
	start := time.Now().UnixNano()

	for i := 0; i < 1; i++ {
		body := search.Get(i)
		fmt.Println(body)
	}

	fmt.Println("直列実行速度 秒", lap(start))
}

func lap(start int64) float64  {
	return float64(time.Now().UnixNano() - start) / 1000000000
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
	}
	fmt.Println("直列実行速度 秒", float64(time.Now().UnixNano() - start) / 1000000000)
}