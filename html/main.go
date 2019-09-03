package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	node, err := html.Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(node)
}
