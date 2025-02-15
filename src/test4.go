package main

import (
	"fmt"
	"sync"
)

func producer(ch chan int, i int) {
	ch <- i
}

func consumer(ch chan int, wg *sync.WaitGroup, wg2 *sync.WaitGroup) {
	for v := range ch {
		fmt.Println(v)
		wg.Done()
	}

	fmt.Println("consumer finish")
	wg2.Done()
}

func main() {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	wg2.Add(1)

	ch := make(chan int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(ch, i)
	}

	go consumer(ch, &wg, &wg2)
	wg.Wait()
	close(ch)
	wg2.Wait()
}