package main

import (
	"fmt"
	"log"
	"runtime"
	"unsafe"
)

func main() {
	fmt.Println(runtime.GOOS,runtime.GOARCH)
	//log.Fatal("heheh")
	log.Println("vvvv")
	fmt.Println(unsafe.Sizeof(int16(0)))
}
