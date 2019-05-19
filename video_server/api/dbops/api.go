package dbops

func AddUserCredential(loginName string,pwd string)  error{
	stmtIns,err:=dbConn.Prepare("insert into user(login_name,pwd) values(?,?)")
	if err!=nil {
		return  err
	}
	stmtIns.Exec(loginName,pwd)
	stmtIns.Close()
	return nil
}

func getUserCredential(loginName string)  (string,error){
	stmtSel,err:=dbConn.Prepare("select pwd from  user where login_name=?")
	if err!=nil {
		return  "",err
	}
	var pwd string
	stmtSel.QueryRow(loginName).Scan(&pwd)
	stmtSel.Close()
	return pwd,nil
}




