package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//得到某一个目录下面的os.FileInfo，使用一个channel来控制能够同时打开的文件数量

func dirents(dir string) []os.FileInfo {
	defer func() {
		<-nfilesCh
	}()

	select {
		case nfilesCh<- struct{}{}:
		case <-done:
			return nil
	}


	f,err:=ioutil.ReadDir(dir)

	if err!=nil {
		fmt.Fprintf(os.Stderr,"du1:%s",err)
		return nil
	}
	return f
}

func walkDir(dir string,filesizes chan<-int64)  {
	//这一句话很关键，在每一次遍历一个目录完成后将wg的数量减少一个
	defer wg.Done()
	if cancelled(){
		return
	}
	//遍历某一个目录
	for _,v:=range dirents(dir) {
		if v.IsDir() {
			wg.Add(1)
			//继续递归遍历
			//fmt.Println("vname:",v.Name())
			subDir:=filepath.Join(dir,v.Name())
			go walkDir(subDir,filesizes)
			//fmt.Printf("%#v",v)
		}else{
			//fmt.Printf("%s file,%.5f size\n",filepath.Join(dir,v.Name()),float64(v.Size()))
			filesizes<-v.Size()
		}
	}
}

func printTrack(nfiles int,sum int64)  {
	fmt.Printf("%d files,total size:%.5fMB\n",nfiles,float64(sum)/1e6)
}

func cancelled()  bool{
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//同步多个goroutine
var wg sync.WaitGroup
//用来控制同时可以打开的目录数量
var nfilesCh chan struct{}
var done chan string
func main()  {
	flag.Parse()
	roots:=flag.Args()
	nfilesCh=make(chan  struct{},1) //最多只能同时打开10个文件
	done=make(chan string)
	fmt.Println("roots:",roots)
	if len(roots)==0 {
		roots=[]string{"."}
	}
	fmt.Println(roots)
	filesizes:=make(chan int64)
	go func() {
		os.Stdin.Read(make([]byte,1))
		close(done)
	}()
	go func() {
		for _,v:=range roots {
			wg.Add(1)
			//fmt.Println("vvvvvvvvv:",v)
			go walkDir(v,filesizes)
		}
	}()

	go func() {
		wg.Wait()
		close(filesizes)
	}()

	var sum int64
	var nfiles int

	tick:=time.Tick(200*time.Millisecond)
	loopb1:
	for {
		select {
				case <-done:
					fmt.Println("done")
					//先要排空filesizes
					for x:=range filesizes {
						fmt.Println("文件大小：",x)
					}
					return
				case <-tick:
					printTrack(nfiles,sum)
				case v,ok:=<-filesizes:
					if !ok {
						break loopb1
					}
					sum+=v
					nfiles++

		}
	}

	for x:=range filesizes {
		nfiles++
		sum+=x
	}
}