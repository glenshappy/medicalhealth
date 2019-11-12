package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

func main() {
	var NodeBits uint
	NodeBits= 10
	aa:= -1 ^ (-1 << NodeBits)
	bb:=2 ^ 3
	fmt.Println(aa,bb)
	n, err := snowflake.NewNode(1)
	snowflake.Epoch = 1288834974999
	if err != nil {
		println(err)
		os.Exit(1)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id", id)
		fmt.Println("node: ", id.Node(), "step: ", id.Step(), "time: ", id.Time(), "\n")
	}
}