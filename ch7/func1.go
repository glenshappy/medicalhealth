package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"
)

type Streamer interface {
	RunningTime() time.Duration
	Format() string
}

type Person struct {
	Name string `json:"name"`
	Age  int16  `json:"age"`
}

type Job []string

func (p Person) RunningTime() time.Duration {
	return 10 * time.Second
}

func (p Person) Format() string {
	return "person"
}

func (j Job) RunningTime() time.Duration {
	return 20 * time.Second
}

func (j Job) Format() string {
	s := ""
	for _, v := range j {
		s += v
	}
	return s
}

func GetFormat(s Streamer) string {
	t := time.Now().UnixNano()
	fmt.Println(t)
	current := strconv.FormatInt(t, 10)
	return current + s.Format()
}

func main() {
	p := Person{"wanjn", 30}
	j := Job{"wanjn", "wanxh", "wancy"}
	rs1 := GetFormat(p)
	rs2 := GetFormat(j)
	fmt.Println(rs1)
	fmt.Println(rs2)
	p2 := []byte{0, 1, 2, 3, 4}
	var w io.Writer = new(bytes.Buffer)
	w.Write(p2)
	//var in1 interface{}
	//in1=[]string{"aa","bb"}
	//in1=222
	//in1="sjsjsjsjs22***我们都是好孩子"
	//fmt.Println(in1)
	//
}
