package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Hello world")
	runtime.GC()
}
