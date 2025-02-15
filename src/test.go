package main

import (
	"fmt"
	"strconv"
)

func main() {

	var hoge [2]int = [2]int{2, 8}
	//append(hoge, 3)
	fmt.Println(hoge)
	//hoge := "\\u6771\\u4eac"
	//unquote, _ := strconv.Unquote(`"` + hoge + `"`)
	//fmt.Println(unquote)

	//hoho := []byte("東京")
	//fmt.Println(hoho)
	//
	//fmt.Println(string(hoho))
	//

	poyo := []byte{91, 34, 92, 117, 54, 55, 55, 49, 92, 117, 52, 101, 97, 99, 34, 93, 125, 125, 125, 93, 125, 125}
	fmt.Println(poyo)
	fmt.Println(string(poyo))

	fuga := []byte("\"")
	fmt.Println(fuga)

	mura := []byte{92, 117, 54, 55, 55, 49, 92, 117, 52, 101, 97, 99}
	fmt.Println(mura)
	fmt.Println(string(mura))
	test, _ := strconv.Unquote(`"` + string(mura) + `"`)
	fmt.Println(test)

	t := "\\u6771\\u4eac0"
	converted, _ := strconv.Unquote(`"` + t + `"`)
	fmt.Println(converted)

	de := `{"took":27,"timed_out":false,"_shards":{"total":1,"successful":1,"skipped":0,"failed":0},"hits":{"total":{"value":1,"relation":"eq"},"max_score":1.0,"hits":[{"_index":"search_job","_type":"_doc","_id":"0","_score":1.0,"_source":{"pref":"\u6771\u4eac0","employment":"full-time","line":{"name":"\u5c71\u624b\u7dda","station":["\u6771\u4eac"]}}}]}}`

	fmt.Println(de)
	deconved, _ := strconv.Unquote(`'` + de + `'`)
	fmt.Println(deconved)

	rune, _, tail, _ := strconv.UnquoteChar("東京", 34)
	fmt.Println(rune)
	fmt.Println(string(rune))
	fmt.Println(tail)




	//fmt.Println("start")
	//ch1, :=, make(chan int)
	//ch2 := make(chan string)
	//go func() {
	//	ch1 <- 100
	//}()
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	ch2 <- "hi"
	//}()
	//
	////v2 := <-ch2
	////fmt.Println(v2)
	////
	////v1 := <-ch1
	////fmt.Println(v1)
	//
	//select {
	//	case v1 := <-ch1:
	//		fmt.Println(v1)
	//	case v2 := <-ch2:
	//		fmt.Println(v2)
	//}
	//
	//select {
	//	case v1 := <-ch1:
	//		fmt.Println(v1)
	//	case v2 := <-ch2:
	//		fmt.Println(v2)
	//}
	//
	//
	//time.Sleep(2 * time.Second)
}