package main

import (
	"fmt"
	"math"
)

const e  = .2828282
const Avogadro  = 6.02214
const Planck  = 6.626e-34

func main(){
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxInt8)
	fmt.Println(1<<2)
	fmt.Println(1<<1)
	var f float32 = 1<<24
	fmt.Println(f)
	fmt.Println(f+1)
	fmt.Println(e)
	fmt.Println(Avogadro)
	fmt.Println(Planck)
	for i:=0;i<11;i++{
		fmt.Printf("x=%d,e^x=%8.3f\n",i,math.Exp(float64(i)))
	}
	var z float64
	fmt.Println(z)
	var nan = math.NaN()
	fmt.Println(nan==nan)
	fmt.Println(math.Sqrt(-1))
}
