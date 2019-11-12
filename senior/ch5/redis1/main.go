package main

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Age int

type Bird interface {
	 Goods(string) string
	 SayHi(int) int
}

func (s *Student) Goods(Name string)  string{
	return Name+"goods"
}

func (s *Student) SayHi(Num int)  int{
	return s.Age+100+Num
}

func toGo(b Bird,name string,num int)  {
	println(b.Goods(name))
	println(b.SayHi(num))
}

func (a Age)Goods(name string)  string{
	return fmt.Sprintf("%s---%d",name,a)
}

func (a Age) SayHi(num int) int{
	return int(a)+num
}

func (s *Student)GetName()  string{
	return s.Name
}


func (s *Student)GetAge()  int{
	return s.Age
}

func (s *Student)PrintStrace()  string{
	b1,err:=json.Marshal(s)
	if err!=nil{
		return ""
	}
	return string(b1)
}

func main() {
	stu1:=&Student{
		"wanjn",
		30,
	}
	fmt.Println("==",stu1,"====")
	var age1 = Age(10)
	toGo(stu1,"wanjn",30)
	toGo(age1,"wanxh",26)
}
