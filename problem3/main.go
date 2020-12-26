package main

import (
	"12-26/problem3/cmd"
	"12-26/problem3/dao"
)

func main() {
	dao.MysqlInit()
	cmd.Entrance()
}