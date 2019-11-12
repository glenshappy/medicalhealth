package deep_equal

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestB(t *testing.T)  {
	rs:=Equal([]int{1,2},[]int{1,2})
	fmt.Println(rs)
	rs1:=Equal(1,2)
	fmt.Println(rs1)
	stu1:=Student{"wanjn",33}
	stu2:=Student{"wanjn",33}
	rs2:=Equal(stu1,stu2)
	fmt.Println(rs2)

}

func TestC(t *testing.T)  {
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "a"}, &link{value: "a"}
	a.tail, b.tail, c.tail = b, c, a
	//fmt.Println(Equal(a, a)) // "true"
	//fmt.Println(Equal(b, b)) // "true"
	//fmt.Println(Equal(c, c)) // "true"
	fmt.Println(Equal(a, b)) // "false"
	//fmt.Println(Equal(c, c)) // "false"
}


func TestA(t *testing.T) {
	a:=strings.Split("a:b:c",":")
	b:=[]string{"a","b","c"}
	var c,d []string=nil,[]string{}
	var e,f map[string]int = nil,make(map[string]int)

	rs:=reflect.DeepEqual(a,b)

	fmt.Println("==========")
	fmt.Println(rs)
	fmt.Println("===========")
	fmt.Println(reflect.DeepEqual(c,d))
	fmt.Println("===========")
	fmt.Println(reflect.DeepEqual(e,f))
	fmt.Println("===========")
}
