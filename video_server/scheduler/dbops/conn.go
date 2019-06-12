package dbops

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(39.106.127.158:3306)/medicalhealth?charset=utf8")
	fmt.Println("数据库连接了一次")
	if err != nil {
		panic(err.Error())
	}
}