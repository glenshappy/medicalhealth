package sessions

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMain(m *testing.M) {

	m.Run()
}

func TestIsSessionExpired(t *testing.T) {
	fmt.Println(time.Now().UnixNano() / 1000000000)
	loginName, err := IsSessionExpired("sjsjsjsj22")
	if err == true {
		t.Errorf("过期或者不存在session")
	}
	fmt.Println(loginName)
}

func TestGenerateNewSessionId(t *testing.T) {
	sid, err := GenerateNewSessionId("wanjn")
	if err != nil {
		t.Errorf("过期的订单：%v", err)
	}
	fmt.Println(sid)
}

func TestLoadSessionsFromDB(t *testing.T) {
	var sessionMap *sync.Map
	sessionMap = &sync.Map{}
	sessionMap.Store("ss", "wanjn")
	fmt.Println(sessionMap.Load("ss"))
	//LoadSessionsFromDB()
}
