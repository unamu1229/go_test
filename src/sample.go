package main

import (
	"fmt"
	"time"
)

func main(){

	sum := make(chan int)
	go func(sum chan int) {
		c := 0
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			c += 1
		}
		sum <- c
	}(sum)

	concat := make(chan string)
	go func(concat chan string) {
		s := ""
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			s = s + "a"
		}
		concat <- s
	}(concat)

	fmt.Println(<-sum)
	fmt.Println(<-concat)
}
