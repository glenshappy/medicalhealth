package dbops

import "testing"

func TestMain(m *testing.M)  {

	m.Run()
}

func TestAddUserCredential(t *testing.T) {
	AddUserCredential("wjn","123")
	if err!=nil{
		t.Errorf("插入失败：%v",err.Error())
	}
}

