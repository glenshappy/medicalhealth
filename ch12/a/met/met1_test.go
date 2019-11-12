package met

import (
	"testing"
	"fmt"
)

func TestMee1(t *testing.T) {
	v:= &Student{
		"wanjn",
		30,
	}
	rs,_:=v.MarshalStudent()
	fmt.Println(rs)
	printMethods(v)
}
