package dbops

import (
	"database/sql"
	"medicalhealth/video_server/api/defs"
)

func GetUserInfo(loginName string)  (*defs.User,error){
	stmtSel,err:=dbConn.Prepare("select id,login_name,pwd,nickname,photo from  user where login_name=?")
	if err!=nil {
		return  &defs.User{},err
	}
	var id int32
	var pwd string
	var loginNameDB string
	var nickname string
	var photo string
	err=stmtSel.QueryRow(loginName).Scan(&id,&loginNameDB,&pwd,&nickname,&photo)
	if err!=nil || err==sql.ErrNoRows {
		return &defs.User{},err
	}
	defer stmtSel.Close()
	return &defs.User{Id:id,LoginName:loginNameDB,Pwd:pwd,Nickname:nickname,Photo:photo},nil
}
