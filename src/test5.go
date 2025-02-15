package main

import "fmt"

func producer(first chan int){
	for i := 0; i < 10; i++ {
		first <- i
	}

	close(first)
}

func secondCalc(first <-chan int, second chan<- int) {
	for v := range first {
		second <- v * 2
	}
	close(second)
}

func thirdCalc(second <-chan int, third chan<- int) {
	for v := range second {
		third <- v * 3
	}
	close(third)
}

func main()  {
	first := make(chan int)
	go producer(first)
	second := make(chan int)
	go secondCalc(first, second)
	third := make(chan int)
	go thirdCalc(second, third)
	for v := range third {
		fmt.Println(v)
	}
}
