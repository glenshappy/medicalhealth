package main

import (
	"fmt"
	"time"
)

const day = 24 * time.Hour

func main() {
	fmt.Println(day.Seconds())
}
