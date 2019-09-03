package dbops

import (
	"database/sql"
	"medicalhealth/video_server/api/defs"
	"medicalhealth/video_server/api/utils"
	"strconv"
	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into user(login_name,pwd) values(?,?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtSel, err := dbConn.Prepare("select pwd from  user where login_name=?")
	if err != nil {
		return "", err
	}
	var pwd string
	err = stmtSel.QueryRow(loginName).Scan(&pwd)
	if err != nil || err == sql.ErrNoRows {
		return "", err
	}
	defer stmtSel.Close()
	return pwd, nil
}

func DelUserCredential(loginName string, pwd string) (bool, error) {
	stmtDel, err := dbConn.Prepare("delete  from user where login_name=? and pwd=?")
	if err != nil {
		return false, err
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return true, nil
}

func AddVideoInfo(userId int32, name string) (defs.VideoInfo, error) {
	//换行请用``而不是""
	stmtIns, err := dbConn.Prepare(`insert into 
		video_info(id,user_id,name,display_ctime,create_time) values(?,?,?,?)`)
	if err != nil {
		return defs.VideoInfo{}, err
	}

	vid, err := utils.NewUUID()
	if err != nil {
		return defs.VideoInfo{}, err
	}
	var createTime = time.Now().Format("2006 1 2 15:04:05")
	stmtIns.Exec(vid, userId, name, createTime, createTime)
	defer stmtIns.Close()
	res := defs.VideoInfo{UserId: userId, Name: name, DisplayCtime: name}
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtSel, err := dbConn.Prepare("select id,user_id,name,display_ctime from  video_info where id=?")
	if err != nil {
		return nil, err
	}
	var id string
	var userId int32
	var name string
	var displayCtime string
	err = stmtSel.QueryRow(vid).Scan(&id, &userId, &name, &displayCtime)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	defer stmtSel.Close()
	return &defs.VideoInfo{UserId: userId, Name: name, DisplayCtime: displayCtime}, nil
}

func DelVideoInfo(vid string) (bool, error) {
	stmtDel, err := dbConn.Prepare("delete  from video_info where id=?")
	if err != nil {
		return false, err
	}
	stmtDel.Exec(vid)
	stmtDel.Close()
	return true, nil
}

func AddComment(videoInfoId string, userId int32, content string) (string, error) {
	//换行请用``而不是""
	stmtIns, err := dbConn.Prepare(`insert into 
		comment(id,video_info_id,user_id,content,time) values(?,?,?,?)`)
	if err != nil {
		return "", err
	}

	cid, err := utils.NewUUID()
	if err != nil {
		return "", err
	}
	var createTime = time.Now().Format("2006 1 2 15:04:05")
	rs, err := stmtIns.Exec(cid, videoInfoId, userId, content, createTime)
	if err != nil {
		return "", nil
	}
	defer stmtIns.Close()
	strInt64OfId, _ := rs.LastInsertId()
	strInt64 := strconv.FormatInt(strInt64OfId, 10)

	return strInt64, nil
}

func GetComment(cid string) (*defs.Comment, error) {
	stmtSel, err := dbConn.Prepare("select id,video_info_id,user_id,content,`time` from  comment where id=?")
	if err != nil {
		return nil, err
	}
	var id string
	var videoInfoId string
	var userId int32
	var ctime string
	var content string
	err = stmtSel.QueryRow(cid).Scan(&id, &videoInfoId, &userId, &content, &ctime)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	defer stmtSel.Close()
	return &defs.Comment{Id: id, VideoInfoId: videoInfoId, UserId: userId, Content: content, Time: ctime}, nil
}

func ListComments(page int32, pageSize int32) (defs.ListComments, error) {
	offset := (page - 1) * pageSize
	stmtSel, err := dbConn.Prepare("select id,video_info_id,user_id,content,`time` from  comment limit ?,?")
	if err != nil {
		return nil, err
	}
	var rsArr defs.ListComments
	var id string
	var videoInfoId string
	var userId int32
	var ctime string
	var content string
	rows, err := stmtSel.Query(offset, pageSize)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	for rows.Next() != false {
		rows.Scan(&id, &videoInfoId, &userId, &ctime, &content)
		rsArr = append(rsArr, defs.Comment{Id: id, VideoInfoId: videoInfoId, UserId: userId, Content: content, Time: ctime})
	}
	defer stmtSel.Close()
	return rsArr, nil
}

func DelComment(id string) (bool, error) {
	stmtDel, err := dbConn.Prepare("delete  from comment where id=?")
	if err != nil {
		return false, err
	}
	stmtDel.Exec(id)
	defer stmtDel.Close()
	return true, nil
}
