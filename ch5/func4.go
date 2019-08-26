package main

import (
	"fmt"
)

type bailout struct {

}

func hehe4()  {
	panic(10)
}

func hehe3(name string)  (err error){
	defer func() {
		p:=recover()
		switch p {
		case nil:
		case bailout {}:
			err=fmt.Errorf("multibile error:%v",p)
		default:
			panic(p)
		}
	}()

	switch name {
	case  "wanjn":
		fmt.Println("nothing")
	default:
		panic(bailout{})
	}
	err=nil
	return
}

func main()  {
	defer func() {
		p:=recover()
		if(p!=nil){
			fmt.Println("======")
			fmt.Println(p)
			fmt.Println("==========")
		}
	}()

	hehe4()

	//err:=hehe3("wanjn2")
	//fmt.Println(err)
	//if err!=nil {
	//	fmt.Println(err)
	//}
	//println("还可以走哈哈")
}
