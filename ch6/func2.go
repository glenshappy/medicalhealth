package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func qpl(arr1 []string) (ret [][]string) {
	if len(arr1) == 0 {
		return
	}
	if len(arr1) == 1 {
		ret[0] = arr1
		return
	}
	for k, v := range arr1 {
		left := append(arr1[:k], arr1[k+1:]...)
		leftpl := qpl(left)
		for _, v1 := range leftpl {
			temp := []string{}
			ret[k] = append(temp, v)
			ret[k] = append(ret[k], v1...)
		}
	}
	return
}

func main() {
	var arr1 [4]string
	arr1 = [4]string{
		"cbm",
		"zob",
		"xoc",
		"wod",
	}
	sl1 := arr1[:]
	rs := qpl(sl1)
	fmt.Println(rs)
}
