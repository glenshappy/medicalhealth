package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 2*time.Second, "ducration")

func main() {
	flag.Parse()
	fmt.Printf("sleepling for %v....", *period)
	time.Sleep(*period)
	fmt.Println()
}
