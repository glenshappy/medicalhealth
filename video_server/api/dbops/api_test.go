package dbops

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMain(m *testing.M)  {

	m.Run()
}

func TestAddUserCredential(t *testing.T) {
	AddUserCredential("wjn","123")
	if err!=nil{
		t.Errorf("插入失败：%v",err.Error())
	}
}

func TestListComments(t *testing.T) {
	rsArr,err:=ListComments(2,1)
	if err!=nil {
		t.Errorf("获取分页失败！")
	}
	dataJson,_:= json.Marshal(rsArr)
	//转换为json字符串得用string方法，否则为[]byte类型
	fmt.Println(string(dataJson))
}

