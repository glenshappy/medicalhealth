package main

import "fmt"

type A interface {
	WriteString(string) error
}

type myint int
type myint1 int
func (my myint)WriteString(s string) (err error) {
	return nil
}

func (my myint)m() string {
	return ""
}

func (my myint)m1() string {
	return "llppp"
}
func (my myint1)m1() string {
	return ""
}

type myint3 int
type myint4 int

func main()  {
	var j interface{m1() string}
	j=myint1(3333)
	//注意这里不能断言成myint，因为j的具体类型是myint1的类型
	if jj,ok:=j.(myint1);ok {
		fmt.Println(jj)
	}else{
		fmt.Println(ok)
		fmt.Println("转换失败2")
	}

}
