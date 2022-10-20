package main

/*
#cgo LDFLAGS: ./lib/libdemo.a -lm -ldl
#include <stdlib.h>
#include "./lib/demo.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

func main() {
	err := testPanic()
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("fin...")
}

func testPanic() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("panic: %v", r))
		}
	}()
	C.trying()
	return
}
