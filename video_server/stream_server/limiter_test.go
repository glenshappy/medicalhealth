package main

import "testing"

func TestMain(m *testing.M)  {

	m.Run()
}

func TestGetConn(t *testing.T) {
	cl:=NewConnLimiter(2)
	rs1:=cl.GetConn()
	if rs1!=true {
		t.Errorf("获取失败1")
	}
	rs2:=cl.GetConn()
	if rs2!=true {
		t.Errorf("获取失败2")
	}
	//rs3:=cl.GetConn()
	//if rs3!=true {
	//	t.Errorf("获取失败3")
	//}
	//rs4:=cl.GetConn()
	//if rs4!=true {
	//	t.Errorf("获取失败4")
	//}
}
func TestReleaseConn(t *testing.T)  {
	cl:=NewConnLimiter(2)
	cl.ReleaseConn()
	cl.ReleaseConn()
}
