package main

import (
	"fmt"
	"sync"
)

func one(x *int)  {
	*x = 77
}

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}

}

func main()  {
	//c := make([]int, 5)
	c := make([]int, 0, 5)
	fmt.Printf("len=%d cap=%d value=%v", len(c), cap(c), c)
	//for i := 0; i < 5; i++ {
	//	c = append(c, 1)
	//	fmt.Println(c)
	//}
	//fmt.Println(c)

}
