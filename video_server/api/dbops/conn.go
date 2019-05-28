package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	dbConn *sql.DB
	err error
)

func init() (){
	fmt.Printf("此时未打开数据库，开始获取一个handler：%v\n",time.Now().String())
	dbConn,err =sql.Open("mysql","root:root@tcp(39.106.127.158:3306)/medicalhealth?charset=utf8")
	dbConn.SetMaxOpenConns(500)
	dbConn.SetMaxIdleConns(200)
	if err!=nil {
		panic(err.Error())
	}
	dbConn.Ping()
}
