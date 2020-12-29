package main

import (
	"fmt"

	"search"
)

func main() {
	body := search.Get()
	fmt.Println(body)
}