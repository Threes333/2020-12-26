package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB
func MysqlInit() {
	var err error
	dsn := "root:qazpl.123456@tcp(127.0.0.1:3306)/message?charset=utf8mb4&parseTime=true&loc=Local"
	DB,err = sql.Open("mysql",dsn)
	if err != nil {
		panic("mysql init failed")
	}
	if err = DB.Ping();err != nil{
		panic("mysql init failed")
	}
	DB.SetMaxOpenConns(1000)
}