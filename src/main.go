package main

import (
	"fmt"
	"search"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		search.Make(i)
		//search.Get()
	}
	//body := search.Get()
	//fmt.Println(body)
	fmt.Println("直列実行速度 秒", float64(time.Now().UnixNano() - start) / 1000000000)
}