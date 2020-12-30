package main

import (
	"fmt"
	"time"

	"search"
)

func main() {
	start := time.Now().UnixNano()
	body := search.Get()
	fmt.Println(body)
	fmt.Println("実行速度 秒", float64(time.Now().UnixNano() - start) / 1000000000)
}