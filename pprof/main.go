package main

import (
	"fmt"
	_ "net/http/pprof"
)
//var data []string
//
//func init() {
//	data = make([]string, 0)
//}

type Node struct {
	Next *Node
}

func main() {
	//ch := make(chan struct{}, 10)
	//go func() {
	//	for {
	//		go func() {
	//			fmt.Println("SSS")
	//			<-ch
	//		}()
	//		time.Sleep(time.Millisecond * 100)
	//	}
	//}()

	//x := 100
	//test := make([]int, x)
	//fmt.Println(test)
	//
	_ = func() int {
		x := 1
		return x
	}()

	_ = func() *int {
		x := new(int)
		return x
	}

	m := map[string]string{
		"a":"a",
		"b":"b",
		"c":"c",
		"d":"d",
		"e":"e",
		"f":"f",
		"g":"g",
		"h":"h",
		"i":"i",
	}
	for index, key := range m {
		fmt.Println(index, key)
	}
}