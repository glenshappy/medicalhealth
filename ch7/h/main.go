package main

import (
	"fmt"
	"io"
	"net/http"
)
//写入字符串
func writeString(w io.Writer,s string)  (n int,err error){
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	//由于断言类型T是一个接口类型，所以w的动态类型必须要实现这个接口才可以断言成功；
	//这个代码请看server.go里面的response类型
	//反之、如果T是一个非接口类型，那么它必须实现w的接口详见：https://colobu.com/2016/06/21/dive-into-go-6/#%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80_type_assertion
	if sw,ok:=w.(stringWriter);ok {
		//fmt.Printf("%#v====%T",w,w)
		//fmt.Println("why not?",ok)
		sw.WriteString(s)
	}
	return w.Write([]byte(s))
}

func writeHeader(w io.Writer,contentType string)  error{
	if _,err:=w.Write([]byte("Content-Type:"));err!=nil {
		return err
	}
	if _,err:=w.Write([]byte(contentType));err!=nil{
		return err
	}
	return nil
}

func price(rw http.ResponseWriter,r *http.Request)  {
	rw.Header()
	itemQ:=r.URL.Query()
	if item,ok:= itemQ["item"];ok {
		c:="TYPE_C:"+item[0]
		err:=writeHeader(rw,c)
		if err!=nil {
			panic(err)
		}
	}
	_,err:=writeString(rw,"what？")
	if err!=nil {
		panic(err)
	}
}


type stringWriter1 interface {
	writeString1(string) (n int, err error)
}

type hehe int

func (h hehe)writeString1(_ string) (n int,err error) {
	return 4,nil
}

func test333(a interface{})  {
	_,ok:=a.(stringWriter1)
	fmt.Println("sjsjsjsjsj")
	fmt.Println(ok)
}

func main()  {
	a:=hehe(22)
	test333(a)

	http.HandleFunc("/test11",price)
	http.ListenAndServe("localhost:9080",nil)
}
