package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle(n int) []int {
	rand.Seed(time.Now().UnixNano())
	return rand.Perm(n)
}

func main() {
	a:=shuffle(10)
	fmt.Println(a)
}
