package met

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Student struct {
	Name string `json:"name,omitempty"`
	Age int `json:"age"`
}

func (s *Student)GetName(s1 string)  string{
	return s.Name
}

func (s *Student)GetAge()  int{
	return s.Age
}

func (s *Student)MarshalStudent()  (string,error){
	sByte,err:=json.Marshal(s)
	if err!=nil {
		return "",err
	}
	return string(sByte),nil
}

func printMethods(v interface{}) {
	uv:=strings.Split("ab_cd_ef_34","_")
	for _,v:=range uv {
		uv=append(uv,v)
	}
	fmt.Println(uv)
	fmt.Println(strings.TrimPrefix("hello_ kitty","hello"))
	val:=reflect.ValueOf(v)
	t:=val.Type()
	fmt.Println("方法数量：",val.NumMethod(),t)
	fmt.Printf("======%#v====\n",t.Elem())
	for i:=0;i<val.NumMethod();i++ {
		methType:=val.Method(i).Type()
		fmt.Printf("func (%s) ---%s----%s\n", t, t.Method(i).Name,
			methType.String(), )
	}
}
