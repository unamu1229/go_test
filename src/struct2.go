package main

import "fmt"

type Hoge struct {
	Y int
	x string
}

func (hoge *Hoge) dubble() int {
	hoge.Y = hoge.Y * hoge.Y
	return hoge.Y
}

func main() {

	hoge := Hoge{
		Y: 10,
		x: "lsls",
	}

	fmt.Println(hoge.dubble())

	fmt.Println("hoge")

}
