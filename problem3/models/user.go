package models

import (
	"12-26/problem3/dao"
	"database/sql"
	"log"
)
type User struct {
	Id       string  `json:"id" form:"id"`
	Username string  `json:"Username" form:"Username"`
	Password string  `json:"password" form:"password"`
	Count    float64 `json:"count" form:"count"`
}
//res = 0 代表注册无问题; res = 1 代表Id已被注册; res = 2 代表用户名已被注册; res = 3 代表Id和用户名都已被注册
func RegisterUser(user *User)(res int){
	res = 0
	var flag interface{}
	sqlstr := "Select * from user where id = ?"
	stmt , _ := dao.DB.Prepare(sqlstr)
	err := stmt.QueryRow(user.Id).Scan(flag)
	if err != sql.ErrNoRows{
		res += 1
	}
	sqlstr = "Select * from user where Username = ?"
	_, err = dao.DB.Prepare(sqlstr)
	stmt , _ = dao.DB.Prepare(sqlstr)
	err = stmt.QueryRow(user.Username).Scan(flag)
	if err != sql.ErrNoRows{
		res += 2
		return res
	}
	sqlstr = "INSERT INTO user (id,Username,password) values(?,?,?)"
	stmt, _ = dao.DB.Prepare(sqlstr)
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Username, user.Password)
	if err != nil{
		log.Println("插入用户数据失败!",err)
	}
	return res
}
//res = 0 代表登录无问题; res = 1 代表用户不存在; res = 2 代表密码输入错误
func LoginUser(user *User) (res int,username string) {
	sqlstr := "Select * from user where id = ?"
	stmt, _ := dao.DB.Prepare(sqlstr)
	defer stmt.Close()
	row := stmt.QueryRow(user.Id)
	var loginuser User
	err := row.Scan(&loginuser.Id, &loginuser.Username, &loginuser.Password)
	if err != nil{
		log.Println("用户不存在!")
		res = 1
		return res,""
	}
	if loginuser.Password != user.Password{
		log.Println("密码输入错误!")
		res = 2
		return res,""
	}
	res = 0
	return res,loginuser.Username
}
