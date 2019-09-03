package main

import (
	"fmt"
	"sort"
)

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	r int
}
type cs []Circle

func (p cs) Len() int           { return len(p) }
func (p cs) Less(i, j int) bool { return p[i].r < p[j].r }
func (p cs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Teacher struct {
	s1 sort.Interface
}

func main() {
	values := []int{1, 4, 3, 1, 91919, 181, 201, 56}
	valuesT := sort.IntSlice(values)
	t1 := Teacher{
		valuesT,
	}
	sort.Sort(t1.s1)
	fmt.Println(t1)
	sort.Sort(sort.Reverse(t1.s1))
	fmt.Println(t1)
	//values:=[]int{1,4,3,1,91919,181,201,56}
	//sort.Sort(sort.Reverse(sort.IntSlice(values)))
	//fmt.Println(values)
	//c:=Circle{Point{1,2},10}
	//fmt.Println(c.Point,c.X,c.Y)
	//cs1:=cs{
	//	Circle{Point{1,2},3},
	//	Circle{Point{2,3},4},
	//	Circle{Point{3,4},5},
	//	Circle{Point{100,101},102},
	//	Circle{Point{97,98},99},
	//	Circle{Point{87,88},89},
	//}
	//sort.Sort(cs1)
	//fmt.Println(cs1)
	////降序排列
	//sort.Sort(sort.Reverse(cs1))
	//fmt.Print(cs1)
}
