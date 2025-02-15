package main

import (
	"fmt"
	"io"
	"os"
)

func LoggingSettings(logFile string) {
	opendLogfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, opendLogfile)
	//log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	//log.SetOutput(multiLogFile)

	result, _ := fmt.Fprint(multiLogFile, "abcdef\n")
	fmt.Println(result)
}


func fuga() (string, string) {
	return "dsdsds", "funefune"
}

type Vertex struct {
	 x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v Vertex) Plus() int {
	return v.x + v.y
}

func New(x, y int) Vertex {
	return Vertex{x, y}
}

func (v Vertex) String() string {
	//return "X is " + strconv.Itoa(v.x) + "! Y is " + strconv.Itoa(v.y) + "!"

	return  fmt.Sprintf("X is %d! Y is %d!", v.x, v.y)
}

func main()  {

	v := Vertex{3, 4}
	fmt.Println(v)
	//fmt.Println(v.Plus())

	//hoge := "tetete"
	//fuga := &hoge
	//
	//fmt.Println(fuga == &hoge)
	//
	//poyo := "tetete"
	//fmt.Println(poyo == hoge)

	//hoge := New(5, 6)
	//fmt.Println(hoge.Area())
	// j 300

	// j 10

	//var hoge []int
	//var p *[]int = &hoge
	//fmt.Println(p)

	//var hoge []int
	//var hoge [9]int
	//fmt.Println(hoge)

	//fmt.Println(*new(int))
	//
	//s := make([]int, 0)
	//fmt.Println(s)

	//var p *int = new(int)
	//var p *int
	//p = new(int)
	//*p = 88
	//fmt.Println(*p)

	//l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	//var minV int
	//for _, v := range l {
	//	if minV == 0 {
	//		minV = v
	//	}
	//	if minV < v {
	//		minV = v
	//	}
	//}
	//fmt.Println(minV)
	//
	//m := map[string]int{
	//	"apple":  200,
	//	"banana": 300,
	//	"grapes": 150,
	//	"orange": 80,
	//	"papaya": 500,
	//	"kiwi":   90,
	//}
	//
	//var sumV int
	//for _, v := range m {
	//	sumV += v
	//}
	//
	//fmt.Println(sumV)

	//hoge := "sss"
	//fmt.Println(hoge)
	//fune, hoge := fuga()
	//fmt.Println(hoge)
	//fmt.Println(fune)

	//if desu := "dddd"; desu != "gege"{
	//	fmt.Println(desu)
	//}

	//fmt.Println(desu)


	//LoggingSettings("hoge.txt")
	//log.Println("dddd")


	//hoge := 888
	////defer fmt.Println(hoge)
	//defer pointer(&hoge)
	//hoge = 9999

	//c := make([]int, 5)

	//c := make([]int, 0, 5)
	//
	//for i := 0; i < 5; i++ {
	//	c = append(c, i)
	//}
	//
	//fmt.Println(c)


	//hoge := 8
	//
	//func () {
	//	fmt.Println(hoge)
	//}()

	//f := 1.11
	//fmt.Println(int(f))
	//
	//// [5, 6]
	//
	//m := map[string]int{
	//	"Mike": 20,
	//	"Nancy":24,
	//	"Messi":30,
	//}
	//fmt.Printf("%T %v", m, m)
}

func pointer(hoge *int) {
	fmt.Println(*hoge)
}