package deep_equal

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type comparision struct {
	x,y unsafe.Pointer
	reflect.Type
}

func Equal(x,y interface{})  bool{
	seen:=make(map[comparision]bool)
	return equal(reflect.ValueOf(x),reflect.ValueOf(y),seen)
}

func equal(x,y reflect.Value,seen map[comparision]bool)  bool{
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type()!=y.Type() {
		return false
	}
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparision{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		fmt.Println(x,y)
		fmt.Println("vjsjjsjsjssj")
		seen[c] = true
	}
	switch x.Kind() {
	case reflect.Struct:
		for i:=0;i<x.NumField();i++ {
			if !equal(x.Field(i),y.Field(i),seen) {

				return false
			}
		}
		return true
	case reflect.Bool:
		return x.Bool()==y.Bool()
	case reflect.Int,reflect.Int16,reflect.Int32,reflect.Int8,reflect.Int64:
		//fmt.Printf("%#v-------%#v------%#v\n",x.Int(),y.Int(),x.Int()==y.Int())
		return x.Int()==y.Int()
	case reflect.String:
		return x.String()==y.String()
	case reflect.Chan,reflect.Func,reflect.UnsafePointer:
		return x.Pointer()==y.Pointer()
	case reflect.Slice,reflect.Array:
		//比较每一个原生
		if x.Len()!=y.Len() {
			return false
		}
		for i:=0;i<x.Len();i++ {
			if !equal(x.Index(i),y.Index(i),seen) {
				return  false
			}
		}
		return true
	case reflect.Ptr,reflect.Interface:
		return equal(x.Elem(),y.Elem(),seen)
	}
	panic("unreachable")
}