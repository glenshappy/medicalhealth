package main

import (
	"fmt"
	"medicalhealth/ch8/d/thumbnail"
	"sync"
)

func makeThumbnails6(filenames <-chan string) {
	var wg sync.WaitGroup
	chan1:=make(chan string)
	for v:=range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			s,err:=thumbnail.ImageFile(f)
			if err!=nil {
				panic(err)
			}
			chan1<-s
		}(v)
	}
	go func() {
		wg.Wait()
		close(chan1)
	}()
	for x:=range chan1 {
		fmt.Println(x)
	}
}

func main()  {
	var filesnames chan string
	filesnames=make(chan  string,5)
	for i:=1;i<=5;i++ {
		path:=fmt.Sprintf("D:/imgs/img%d.jpg",i)
		filesnames<-path
	}
	close(filesnames)
	makeThumbnails6(filesnames)
}