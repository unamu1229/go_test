package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() {
		ch1 <- 100
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "hi"
	}()

	//v2 := <-ch2
	//fmt.Println(v2)
	//
	//v1 := <-ch1
	//fmt.Println(v1)

	select {
		case v1 := <-ch1:
			fmt.Println(v1)
		case v2 := <-ch2:
			fmt.Println(v2)
	}

	select {
		case v1 := <-ch1:
			fmt.Println(v1)
		case v2 := <-ch2:
			fmt.Println(v2)
	}


	time.Sleep(2 * time.Second)
}