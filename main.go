package main

import (
	"github.com/Adebusy/powerusyservice/app/driver"
	controllersroute "github.com/Adebusy/powerusyservice/controllersRoute"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DbGorm *gorm.DB
var ErrGorm error

var userRoute = controllersroute.NewIuserController(driver.GetDB())

func main() {
	driver.CreateLog()
	svc := gin.Default()
	svc.GET("/", userRoute.CheckService)
	svc.POST("/createUser", userRoute.CreateNewUser)
	svc.POST("/Login", userRoute.LoginUSer)
	svc.GET("/GetAllUser", userRoute.GetAllUser)
	svc.GET("/CheckEmailWithAuthCode", userRoute.CheckEmailWithAuthCode)
	svc.Run(":8060")
}
