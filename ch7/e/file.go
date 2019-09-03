package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

var ErrNotExist = errors.New("file does not exist")

func isNotExist(err error)  bool{
	fmt.Printf("%#v\n",err)
	if pe,ok:=err.(*os.PathError);ok { //将接口类型转换为具体的类型
		err=pe.Err
	}
	fmt.Printf("%T\n",err)
	fmt.Printf("%T\n",syscall.ENOENT)
	return err==syscall.ENOENT || err==ErrNotExist || err==syscall.ENOTDIR
}

func main()  {
	//var errs = errors.New("file does not exist")
	var _,err= os.Open("/d/test113444.log")
	fmt.Println(isNotExist(err))
}
