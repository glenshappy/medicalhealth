package main

import (
	"errors"
	"fmt"
)

type s uintptr

func main() {
	err := fmt.Errorf("%s", "sjsj29292")
	fmt.Println(errors.New("EOF"), err)
}
