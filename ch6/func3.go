package main

import (
	"fmt"
	"net/url"
)

type P int

func (p P) f() {
	fmt.Println(2222)
}

type Student struct {
	name string
	age  int16
}

func (s *Student) f1() {
	s.name = "wxh"
}

func hehe(yu url.Values) {
	yu.Add("lang", "sjs")
	yu = nil
}

func main() {
	//var o P = 3
	//o.f()
	//
	//sWjn:= Student{
	//	"wanjn",19,
	//}
	//(&sWjn).f1()
	//fmt.Println(sWjn)
	//
	//s2:= &sWjn
	//s2.name="ajajajajaja"
	//fmt.Println(sWjn)

	//s:=make(map[string]int)
	//if s==nil {
	//	fmt.Println("nil")
	//}
	//s["wanjn"] = 222
	//fmt.Printf("%#v\n",s)
	//
	//s1:=s
	//s1["wanjn"]  = 33
	//fmt.Printf("%#v\n",s1)
	//
	//var s2 *Student
	//if s2==nil {
	//	fmt.Println(39)
	//}
	//fmt.Printf("%#v",s2)
	m := url.Values{"lang": {"en", "zh"}}
	//m.Add("lang","ch")
	//fmt.Println(m["lang"])
	//hehe(m)
	//fmt.Println("===========")
	//fmt.Println(m)
	//fmt.Println("============")
	//s:=m
	//s.Add("lang","wh")
	hehe(m)
	s := m
	s.Add("lang", "iiii")
	s = nil
	fmt.Println("=======")
	fmt.Println(m)
	fmt.Println("=======")

}
