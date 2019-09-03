package main

import "fmt"

//注意s的类型不是接口类型，不能用来断言，另外实参也不可以传递[]map[string]int这种类型
//总结一点：形参和实参类型一定要对应，如果是接口一定要实现了才可以传递
func test(s []map[string]interface{})  {
	for _,v:=range s {
		for i1,v1:=range v {
			switch v1.(type) {
			case int:
				fmt.Println("int:",i1,v1)
			default:
				fmt.Println("error")
			}
		}
	}
}
func main()  {
	var u = []map[string]interface{}{
		{"aa":222},
		{"bbb":333},
	}
	test(u)
}
