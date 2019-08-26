package main

import (
	"fmt"
	"math"
)

func test111(bit uint8)  {
	var a uint64
	a=1
	a|=3<<bit
	fmt.Println(a)
}

type Point struct {X,Y int}

func (p Point)  path(q Point) float64 {
	x:=(p.X-q.X)*(p.X-q.X)
	y:=(p.Y-q.Y)*(p.Y-q.Y)
	return math.Sqrt(float64(x+y))
}

type Path []Point

func (p Point) Add(q Point) Point{
	return Point{p.X+q.X,p.Y+q.Y}
}

func (p Point) Sub(q Point)  Point{
	return Point{p.X-q.Y,p.Y-q.Y}
}

func (path Path)TranslateBy(offset Point,add bool)  {
	var op func(p Point,q Point) Point
	if add==true {
		 op=Point.Add
	}else{
		op=Point.Sub
	}
	for i,v:=range path {
		path[i] = op(v,offset)
	}
}

type Counter struct {
	n int
}

type Buffer struct {
	buf []byte
	initial [3]byte
}

func (b *Buffer)Grow(n int)  {
	if b.buf==nil {
		b.buf=b.initial[:3]
	}
	if len(b.buf)+n > cap(b.buf) {
		buf:=make([]byte,len(b.buf),2*cap(b.buf)+n)
		copy(buf,b.buf)
		b.buf = buf
	}
}

func main()  {

	var b Buffer
	b = Buffer{initial:[3]byte{1,2,30}}
	b.Grow(300)
	fmt.Println(b)
	fmt.Println(len(b.buf),cap(b.buf))



	var a = []byte{1,2,3}
	var b1 = []byte{122,251}
	aa:=copy(a,b1)
	fmt.Println(aa,a)


	var c1 []string
	c1=[]string{"a","b","c"}
	var c2 = []string{"1","2"}
	copy(c1,c2)
	fmt.Println(c1)

	var c3=[]byte{1,3,4,5,6,7,8}
	var c4="🐧aa"
	copy(c3,c4)
	fmt.Println(c3)
	//var hehe []Point
	////hehe = make([]Point,3)
	//if nil==hehe {
	//	fmt.Println("uuiiii")
	//}
	//hehe[0] = Point{1,3}
	//fmt.Printf("%#v",hehe)
	//
	//fmt.Println(hehe)
	//var path = Path{
	//	{1,2},
	//	{3,4},
	//	{5,6},
	//}
	//var offset = Point{10,10}
	//path.TranslateBy(offset,false)
	//fmt.Println(path)
	//p:=Point{1,1}
	//q:=Point{4,5}
	//pf:=Point.path
	//
	//sss:=pf(p,q)
	//
	//s:=p.path(q)
	//fmt.Println(s,sss)
}
