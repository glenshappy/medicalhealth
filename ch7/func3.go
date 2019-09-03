package main

import (
	"flag"
	"fmt"
	"github.com/tempconv"
)

//var d = flag.Duration("period",1*time.Second,"period")

var temp = tempconv.CelsiusFlag("temp", 20.0, "temperature")

func main() {
	//var c Celsius = 11
	//fmt.Println(c)
	//flag.Parse()
	//time.Sleep(*d)
	//fmt.Printf("sleep for：%v\n",*d)

	flag.Parse()
	fmt.Printf("sjsjsj:%v", *temp)
}
