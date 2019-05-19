package dbops

import (
	"database/sql"
	"medicalhealth/video_server/api/defs"
	"medicalhealth/video_server/api/utils"
	"time"
)

func AddUserCredential(loginName string,pwd string)  error{
	stmtIns,err:=dbConn.Prepare("insert into user(login_name,pwd) values(?,?)")
	if err!=nil {
		return  err
	}
	stmtIns.Exec(loginName,pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string)  (string,error){
	stmtSel,err:=dbConn.Prepare("select pwd from  user where login_name=?")
	if err!=nil {
		return  "",err
	}
	var pwd string
	err=stmtSel.QueryRow(loginName).Scan(&pwd)
	if err!=nil || err==sql.ErrNoRows {
		return "",err
	}
	defer stmtSel.Close()
	return pwd,nil
}

func DelUserCredential(loginName string,pwd string)  (bool,error){
	stmtDel,err:=dbConn.Prepare("delete  from user where login_name=? and pwd=?")
	if err!=nil {
		return  false, err
	}
	stmtDel.Exec(loginName,pwd)
	stmtDel.Close()
	return  true,nil
}

func AddVideoInfo(userId int32,name string)  (defs.VideoInfo,error){
	//换行请用``而不是""
	stmtIns,err:=dbConn.Prepare(`insert into 
		video_info(id,user_id,name,display_ctime,create_time) values(?,?,?,?)`)
	if err!=nil {
		return  defs.VideoInfo{},err
	}

	vid,err:=utils.NewUUID()
	if err!=nil {
		return  defs.VideoInfo{},err
	}
	var createTime  = time.Now().Format("2006 1 2 15:04:05")
	stmtIns.Exec(vid,userId,name,createTime,createTime)
	defer stmtIns.Close()
	res:=defs.VideoInfo{UserId:userId,Name:name,DisplayCtime:name}
	return res,nil
}

func GetVideoInfo(vid string)  (*defs.VideoInfo,error){
	stmtSel,err:=dbConn.Prepare("select id,user_id,name,display_ctime from  video_info where vid=?")
	if err!=nil {
		return  nil,err
	}
	var id string
	var userId int32
	var name string
	var displayCtime string
	err=stmtSel.QueryRow(vid).Scan(&id,&userId,&name,&displayCtime)
	if err!=nil || err==sql.ErrNoRows {
		return nil,err
	}
	defer stmtSel.Close()
	return &defs.VideoInfo{UserId:userId,Name:name,DisplayCtime:displayCtime},nil
}

func DelVideoInfo(vid string)  (bool,error){
	stmtDel,err:=dbConn.Prepare("delete  from video_info where id=?")
	if err!=nil {
		return  false, err
	}
	stmtDel.Exec(vid)
	stmtDel.Close()
	return  true,nil
}