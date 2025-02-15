package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v map[string] int
	m sync.Mutex
}

func (c *Counter) up(key string) {
	c.m.Lock()
	c.v[key]++
	c.m.Unlock()
}

func main(){

	c := Counter{v: make(map[string]int)}
	//c := make(map[string]int)
	//c := 0
	for i := 0; i < 30; i++ {
		go func () {
			for i := 0; i < 10; i++ {
				c.up("key")
				//c["key"] += 1
				//c += 1
			}
		}()

		go func () {
			for i := 0; i < 10; i++ {
				c.up("key")
				//c["key"] += 1
				//c += 1
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(c, c.v["key"])
}
