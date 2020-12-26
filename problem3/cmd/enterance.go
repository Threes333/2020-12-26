package cmd

import (
	"12-26/problem3/controller"
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	r.GET("/login",controller.LoginUser)
	r.POST("/register",controller.RegisterUser)
	r.POST("/recharge",controller.SelfRecharge)
	r.PUT("/transfer",controller.Transfer)
	r.GET("/record",controller.ViewRecord)
	r.GET("/SelfRecord",controller.CheckSelfRecord)
	_ = r.Run(":8080")
}
