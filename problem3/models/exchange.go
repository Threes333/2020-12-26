package models

import (
	"12-26/problem3/dao"
	"errors"
	"fmt"
	"log"
)

type Affair struct {
	Sender string`json:"sender" form:"sender"`
	Receiver string`json:"receiver" form:"receiver"`
	Count  float64`json:"count" form:"count"`
	Message string`json:"message" form:"message"`
}
func SelfRecharge(affair *Affair) error {
	var flag interface{}
	sqlstr := "select id from user where username = ?"
	stmt, _ := dao.DB.Prepare(sqlstr)
	err := stmt.QueryRow(affair.Receiver).Scan(&flag)
	if err != nil{
		fmt.Println("转入用户不存在")
		return errors.New("转入用户不存在")
	}
	sqlstr = "Update user set count = count + ?"
	stmt, _ = dao.DB.Prepare(sqlstr)
	_, err = stmt.Exec(affair.Count)
	if err != nil{
		fmt.Println("充值失败")
		return	errors.New("充值失败")
	}
	sqlstr = "INSERT INTO affair (sender,receiver,count,message) values (?,?,?,?)"
	stmt, _ = dao.DB.Prepare(sqlstr)
	defer stmt.Close()
	_, err = stmt.Exec(affair.Sender, affair.Receiver, affair.Count, affair.Message)
	if err != nil {
		fmt.Println("插入充值记录失败")
		return	errors.New("插入充值记录失败")
	}
	return nil
}
func Transfer(affair *Affair) error {
	var flag interface{}
	sqlstr := "select id from user where username = ?"
	stmt, _ := dao.DB.Prepare(sqlstr)
	err := stmt.QueryRow(affair.Receiver).Scan(&flag)
	if err != nil{
		fmt.Println("转入用户不存在")
		return errors.New("转入用户不存在")
	}
	var count float64
	sqlstr = "select count from user where username = ?"
	stmt,_ = dao.DB.Prepare(sqlstr)
	err = stmt.QueryRow(affair.Receiver).Scan(&count)
	if err != nil {
		return errors.New("转出用户不存在")
	}
	if count - affair.Count < 0{
		return errors.New("转出用户余额不足")
	}
	tx, _ := dao.DB.Begin()
	sqlstr = "Update user set count = count + ?"
	stmt, _ = tx.Prepare(sqlstr)
	_, err = stmt.Exec(affair.Count)
	if err != nil{
		fmt.Println("充值失败")
		_ = tx.Rollback()
		return errors.New("充值失败")
	}
	sqlstr = "Update user set count = count - ?"
	stmt, _ = tx.Prepare(sqlstr)
	_, err = stmt.Exec(affair.Count)
	if err != nil{
		fmt.Println("充值失败")
		_ = tx.Rollback()
		return errors.New("充值失败")
	}
	sqlstr = "INSERT INTO affair (sender,receiver,count,message) values (?,?,?,?)"
	stmt, _ = dao.DB.Prepare(sqlstr)
	defer stmt.Close()
	_, err = stmt.Exec(affair.Sender, affair.Receiver, affair.Count, affair.Message)
	if err != nil {
		fmt.Println("插入充值记录失败")
		_ = tx.Rollback()
		return errors.New("插入充值记录失败")
	}
	_ = tx.Commit()
	return nil
}

func ViewRecord(msg string) ([]Affair,error) {
	sqlstr := "select sender, count, receiver, message from affair where message like ?"
	stmt, _ := dao.DB.Prepare(sqlstr)
	defer stmt.Close()
	rows, err := stmt.Query(msg)
	if err != nil {
		log.Println(err)
		return nil,errors.New("查询失败")
	}
	affair := make([]Affair,0)
	for rows.Next(){
		var afir Affair
		_ = rows.Scan(&afir.Sender,&afir.Count,&afir.Receiver,&afir.Message)
		affair = append(affair,afir)
	}
	return affair,nil
}
func CheckSelfRecord(username string) ([]Affair,error) {
	affair := make([]Affair,0)
	sqlstr := "select sender,receiver,count,message,create_time from affair where sender = ? or receiver = ? order by create_time"
	stmt, _ := dao.DB.Prepare(sqlstr)
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Println(err)
		return nil,errors.New("查询失败")
	}
	for rows.Next(){
		var afir Affair
		_ = rows.Scan(&afir.Sender,&afir.Count,&afir.Receiver,&afir.Message)
		affair = append(affair,afir)
	}
	return affair,nil
}