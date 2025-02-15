package main

import (
	"fmt"
)

func hoge(c chan int) {
	c <- 1
}

func main() {

	// channel
	fmt.Println("hello again")
	c := make(chan int)
	go hoge(c)
	var result int
	result = <-c
	fmt.Println(result)

	// bufferd channel
	bc := make(chan int, 2)
	go hoge(bc)
	go hoge(bc)
	result = <-bc
	fmt.Println(result)
	result = <-bc
	fmt.Println(result)
	// バッファードチャネルのチャンネル数をこえて取得しようとするので
	// fatal error: all goroutines are asleep - deadlock!
	//result = <-bc
	fmt.Println(result)
}
