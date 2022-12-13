package main

/*
#cgo LDFLAGS: ./lib/libdemo.a -lm -ldl
#include <stdlib.h>
#include "./lib/demo.h"
*/
import "C"
import (
	"fmt"
)

func main() {
	fmt.Println("-----start-------")
	err := testPanic()
	if err != nil {
		fmt.Println("------------------err:", err)
	}
	fmt.Println("----------------------fin...")
}

func testPanic() (err error) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("!!!!!!!!!!!!panic!!!!!!!!!!!!!! ", r)
	//		err = errors.New(fmt.Sprintf("---------------panic: %v", r))
	//	}
	//}()
	C.trying()
	fmt.Println("-----------tried!!!")
	return
}
