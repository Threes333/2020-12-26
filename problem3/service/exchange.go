package service

import (
	"12-26/problem3/models"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func SelfRecharge(c *gin.Context) error {
	var affair models.Affair
	err := c.ShouldBind(&affair)
	if err != nil {
		log.Println(err)
		return errors.New("接受交易数据失败")
	}
	return models.SelfRecharge(&affair)
}
func Transfer(c *gin.Context) error {
	var affair models.Affair
	err := c.ShouldBind(&affair)
	if err != nil {
		log.Println(err)
		return errors.New("接受交易数据失败")
	}
	return models.Transfer(&affair)
}
func ViewRecord(c *gin.Context) (string,error) {
	var message string
	message = c.PostForm("message")
	msg,err := models.ViewRecord(message)
	if err != nil {
		return "", err
	}
	bytes,_ := json.Marshal(msg)
	return	string(bytes),nil
}
func CheckSelfRecord(c *gin.Context) (string,error) {
	var username string
	username = c.PostForm("username")
	record,err := models.CheckSelfRecord(username)
	if err != nil{
		return "",err
	}
	bytes,_ := json.Marshal(record)
	return	string(bytes),nil
}
