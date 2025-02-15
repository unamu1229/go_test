package main

import "fmt"

func goRourtine(s []int, c chan int) {
	for _, v := range s {
		c <- v
	}
	close(c)
}

func main() {
	s := []int{18, 2, 3, 4, 5}
	c := make(chan int)
	go goRourtine(s, c)

	for v := range c {
		fmt.Println(v)
	}
}
