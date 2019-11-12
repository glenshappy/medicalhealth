package refl

import (
	"fmt"
	"reflect"
	"strconv"
)

func formatAtom(v reflect.Value)  string{
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.String:
		return v.String()
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return strconv.FormatInt(v.Int(),10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Chan,reflect.Func,reflect.Ptr,reflect.Map,reflect.Slice:
		return v.Type().String()+"---0x"+strconv.FormatUint(uint64(v.Pointer()),16)
	default:
		return v.Type().String()+"value"
	}
}

func judgeV(path string,v reflect.Value) {
	switch v.Kind() {
	case reflect.Slice:
		for i:=0;i<v.Len();i++ {
			s:=v.Index(i)
			fmt.Printf("%#v----%T\n",s,s)
		}
	case reflect.Struct:
		//反射可以将结构体当做数组一样进行遍历
		for i:=0;i<v.NumField();i++ {
			filepath:=fmt.Sprintf("%s",v.Type().Field(i).Name)
			judgeV(filepath,v.Field(i))
		}
	default:
		rs:=formatAtom(v)
		fmt.Println(rs)
	}
}
