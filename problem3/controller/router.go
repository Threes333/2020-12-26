package controller

import (
	"12-26/problem3/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(c *gin.Context){
	switch service.RegisterUser(c) {
	case 0:
		c.JSON(http.StatusOK,gin.H{
			"code" : "10000",
			"message" : "success register",
		})
	case 1:
		c.JSON(http.StatusOK,gin.H{
			"code" : "20001",
			"message" : "Id已被注册",
		})
	case 2:
		c.JSON(http.StatusOK,gin.H{
			"code" : "20002",
			"message" : "用户名已被注册",
		})
	case 3:
		c.JSON(http.StatusOK,gin.H{
			"code" : "20003",
			"message" : "Id和用户名都已被注册",
		})
	case 4:
		c.JSON(http.StatusOK,gin.H{
			"code" : "20004",
			"message" : "Id或用户名或密码为空",
		})
	}
}

func LoginUser(c *gin.Context){
	switch service.LoginUser(c) {
	case 0:
		c.JSON(http.StatusOK,gin.H{
			"code" : "10000",
			"message" : "success login",
		})
	case 1:
		c.JSON(http.StatusOK,gin.H{
			"code" : "20001",
			"message" : "用户不存在",
		})
	case 2:
		c.JSON(http.StatusOK,gin.H{
			"code" : "20002",
			"message" : "密码输入错误",
		})
	case 3:
		c.JSON(http.StatusOK,gin.H{
			"code" : "20003",
			"message" : "用户名或密码为空",
		})
	}
}
func SelfRecharge(c *gin.Context) {
	err := service.SelfRecharge(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code" : "20000",
			"message" : err,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code" : "10000",
			"message" : "充值成功",
		})
	}
}
func Transfer(c *gin.Context) {
	err := service.Transfer(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code" : "20000",
			"message" : err,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code" : "10000",
			"message" : "转账成功",
		})
	}
}
func ViewRecord(c *gin.Context) {
	message,err := service.ViewRecord(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code" : "20000",
			"message" : err,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code" : "10000",
			"message" : message,
		})
	}
}

func CheckSelfRecord(c *gin.Context) {
	record,err := service.CheckSelfRecord(c)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{
			"code" : "20000",
			"message" : err,
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code" : "10000",
			"message" : record,
		})
	}
}
