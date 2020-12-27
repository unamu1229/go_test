package main

import (
	"fmt" 
	"time"
	"strconv"
)

func sub(count int, hoge *string) {
	fmt.Println("sub" + strconv.Itoa(count) + " is running")
	time.Sleep(time.Second)
	fmt.Println("sub" + strconv.Itoa(count) + " is finished")
	*hoge += "sub" + strconv.Itoa(count)
//	tasks <- "sub" + strconv.Itoa(count)
}

func main() {
	fmt.Println("start")
//	tasks = make(chan string)
	var hoge = ""
	go sub(1, &hoge)
//	hoge2 := <-tasks
	// hoge = ""
	go sub(2, &hoge)
	go sub(3, &hoge)
	go sub(4, &hoge)
	go func() {
		fmt.Println("即時関数")
	}()
	time.Sleep(2 * time.Second)

	fmt.Println(hoge)
}