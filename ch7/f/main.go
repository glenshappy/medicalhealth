package main

import "fmt"

type Var []int

func main()  {
	var x interface{}
	x = Var{1,2,3}
	switch x:=x.(type) {
	case int:
		println(x)
	case []string:
		for _,v:=range x {
			println(v)
		}
	default:
		panic(fmt.Sprintf("unexpected type %T---%v",x,x))
	}
}
