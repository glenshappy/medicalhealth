package main

import "fmt"

type reader11 int
type readerwriter11 int

type student struct {
	name string
}

func (r student)A2(_ string) string {
	return "write_student"
}

func (r student)A1(_ string) string {
	return "read_student"
}

func (r reader11)A1(_ string) string {
	return "s"
}

func (r readerwriter11)A2(_ string) string {
	return "write"
}

func (r readerwriter11)A1(_ string) string {
	return "read"
}




type Reader1 interface {
	A1(s string) (n string)
}

type ReaderWriter1 interface {
	Reader1
	A2(s string) (n string)
}

func main()  {
	//var r1  Reader1
	//r1=reader11(2)
	//
	//r2,ok:=r1.(ReaderWriter1)
	//fmt.Println(r2,ok)

	var r2 ReaderWriter1
	r2=readerwriter11(22)
	r1,_:=r2.(Reader1) //1、大的数据类型一定可以转换成小的数据类型，相当于赋值操作；
	// 2、这里转换成了Reader1的接口类型，代表只可以调用A1方法，A2无法调用；
	//3、r2经过类型断言以后，r1的具体类型依然是r2的具体类型；
	fmt.Println(r1.A1("ss"))

	//var r22 ReaderWriter1
	//var ok bool
	//var stu  = student{"wanjn"}
	//r22,ok=stu.(ReaderWriter1) //为什么不能转换？非接口不能转接口？
	//fmt.Println(r22,ok)

	//var stu  = student{"wanjn"}
	r22,ok:=r2.(student)
	fmt.Println("============")
	fmt.Println(r22,ok)

	//var r3 Reader1
	//fmt.Printf("%T===%T",r1,r3)

	//var w io.Writer
	//w=os.Stdout
	//b,ok:=w.(*bytes.Buffer)
	//
	//rw,ok:=w.(io.ReadWriter)
	//fmt.Println("start==========test===========")
	//fmt.Printf("%#v\n",rw)
	//fmt.Println("end==========test===========")
	//
	//if !ok {
	//	fmt.Printf("%#v",b)
	//	if b==nil {
	//		fmt.Println("what?")
	//	}
	//}
	//if w,ok:=w.(*os.File);ok {
	//	fmt.Printf("%#v",w)
	//}

}
