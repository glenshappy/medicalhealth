package main

import (
	"fmt"
	"sync"
)

var icons map[string]string
func loadIcons()  {
	fmt.Println("搞事情")
	icons=make(map[string]string)
	icons=map[string]string{
		"aaa":"aaa1",
		"bbb":"bbb1",
		"ccc":"ccc1",
	}
}

func Icon(name string)  (string,bool){
	/*
	每一次对Do(loadIcons)的调用都会锁定mutex，并检查boolean变量。在第一次调用
	时，变量的值是false,Do会调用loadIcons，Do会调用loadIcons并将boolean设置为true。
	随后的调用什么都不会做，但是mutex同步会保证

	“loadIcons对内存产生的效果对所有goroutine可见”，用这种方式来使用sync.Once
	的话，我们能够避免在变量被构建完成之前和其他goroutine共享该变量。
	*/
	//loadIconsOnce.Do(loadIcons)//全局保证loadIcons只会被执行一次
	if icons==nil {
		loadIcons()
	}
	v,ok:=icons[name]
	return v,ok
}
var wg sync.WaitGroup
var loadIconsOnce sync.Once
func main()  {
	wg.Add(4)
	go func() {
		if v,ok:=Icon("aaa");ok {
			fmt.Println("携程1：",v)
		}else{
			fmt.Println("携程1：","不存在！")
		}
		wg.Done()
	}()

	go func() {
		if v,ok:=Icon("bbb");ok {
			fmt.Println("携程2：",v)
		}else{
			fmt.Println("携程2：","不存在！")
		}
		wg.Done()
	}()

	go func() {
		if v,ok:=Icon("ccc");ok {
			fmt.Println("携程3：",v)
		}else{
			fmt.Println("携程3：","不存在！")
		}
		wg.Done()
	}()
	go func() {
		if v,ok:=Icon("ddd");ok {
			fmt.Println("携程4：",v)
		}else{
			fmt.Println("携程4：","不存在！")
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("结束")
}
