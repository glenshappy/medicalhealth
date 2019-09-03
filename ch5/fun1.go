package main

import (
	"fmt"
	"os"
)

const (
	AA, BIT0 = 1 << iota, iota - 1

	BB, BIT1

	CC, BIT2

	DD, GG10 = 1, iota

	EE, GG1

	FF, GG1112 = 1 << iota, iota

	GG, GG2
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "line %d:", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func max(args ...int) int {
	var max = 0
	for _, v := range args {
		if v > max {
			max = v
		}
	}
	return max
}

func min(args ...int) int {
	var min = 0
	for _, v := range args {
		if v < min {
			min = v
		}
	}
	return min
}

func hehe() {
	defer println("over")
	var map1 map[string]int
	map1["name"] = 30

}

func hehe1(s string) {
	switch s {
	case "name":
		println(222)
	case "age":
		println(333)
	default:
		panic("error")
	}
}

func main() {
	i := 2
	fmt.Println(i)
	hehe1("sjsjsj")
	hehe()
	var tempdirs = []int{
		1, 2, 3, 4,
	}
	var funcs []func()

	for _, v := range tempdirs {
		funcs = append(funcs, func() {
			fmt.Println(v)
		})
	}

	fmt.Println("================")
	for _, v := range funcs {
		v()
	}
	fmt.Println("=================")

	max := max()
	min := min(22, -100, 23, -232)
	fmt.Println(max, min)
	linenum, format := 12, "error %s--%s:"
	errorf(linenum, format, "wrong args", "nihaode")
	//fmt.Println(AA,BB,CC,DD,EE,FF,GG,BIT0,BIT1,BIT2,GG1,GG2)
	//f:=squares()
	//fmt.Println(f())
	//fmt.Println(f())
	//fmt.Println(f())

	//var f1 func(int) int
	//if f1==nil {
	//	println("f1为空")
	//	f1(3)
	//}

}
