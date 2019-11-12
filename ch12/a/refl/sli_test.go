package refl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSli(t *testing.T) {
	ch1:=make(chan  int,3)
	val1:=reflect.ValueOf(ch1)
	judgeV(val1)
}

func TestNew1(t *testing.T)  {
	fmt.Println(33)
}