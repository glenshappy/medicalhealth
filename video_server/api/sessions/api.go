package sessions

import (
	"database/sql"
	"medicalhealth/video_server/api/defs"
	"medicalhealth/video_server/api/utils"
	"strconv"
	"time"
)

func AddSession(videoInfoId string,userId int32,content string)  (string,error){
	//换行请用``而不是""
	stmtIns,err:=dbConn.Prepare(`insert into 
		comment(id,video_info_id,user_id,content,time) values(?,?,?,?)`)
	if err!=nil {
		return  "",err
	}

	cid,err:=utils.NewUUID()
	if err!=nil {
		return  "",err
	}
	var createTime  = time.Now().Format("2006 1 2 15:04:05")
	rs,err:=stmtIns.Exec(cid,videoInfoId,userId,content,createTime)
	if err!=nil {
		return "",nil
	}
	defer stmtIns.Close()
	strInt64OfId,_:= rs.LastInsertId()
	strInt64:=strconv.FormatInt(strInt64OfId,10)

	return strInt64,nil
}

func GetComment(cid string)  (*defs.Comment,error){
	stmtSel,err:=dbConn.Prepare("select id,video_info_id,user_id,content,`time` from  comment where id=?")
	if err!=nil {
		return  nil,err
	}
	var id string
	var videoInfoId string
	var userId int32
	var ctime string
	var content string
	err=stmtSel.QueryRow(cid).Scan(&id,&videoInfoId,&userId,&content,&ctime)
	if err!=nil || err==sql.ErrNoRows {
		return nil,err
	}
	defer stmtSel.Close()
	return &defs.Comment{Id:id,VideoInfoId:videoInfoId,UserId:userId,Content:content,Time:ctime},nil
}

func ListComments(page int32,pageSize int32)  (defs.ListComments,error){
	offset:=(page-1)*pageSize
	stmtSel,err:=dbConn.Prepare("select id,video_info_id,user_id,content,`time` from  comment limit ?,?")
	if err!=nil {
		return  nil,err
	}
	var rsArr defs.ListComments
	var id string
	var videoInfoId string
	var userId int32
	var ctime string
	var content string
	rows,err:=stmtSel.Query(offset,pageSize)
	if err!=nil || err==sql.ErrNoRows {
		return nil,err
	}
	for rows.Next()!=false {
		rows.Scan(&id,&videoInfoId,&userId,&ctime,&content)
		rsArr = append(rsArr,defs.Comment{Id:id,VideoInfoId:videoInfoId,UserId:userId,Content:content,Time:ctime})
	}
	defer stmtSel.Close()
	return rsArr,nil
}

func DelComment(id string)  (bool,error){
	stmtDel,err:=dbConn.Prepare("delete  from comment where id=?")
	if err!=nil {
		return  false, err
	}
	stmtDel.Exec(id)
	defer stmtDel.Close()
	return  true,nil
}
