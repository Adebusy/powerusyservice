package main

import (
	"github.com/Adebusy/powerusyservice/app/driver"
	controllersroute "github.com/Adebusy/powerusyservice/controllersRoute"
	"github.com/Adebusy/powerusyservice/utilities"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/Adebusy/powerusyservice/docs/powerusyservice" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var DbGorm *gorm.DB
var ErrGorm error

var userRoute = controllersroute.NewIuserController(driver.GetDB())

// @title Powerusy backend service
// @version 1.0
// @description This is a backend web service created for Powerusy operations.
// @termsOfService http://swagger.io/terms/

// @contact.name Alao Ramon, Nabil Abubakar
// @contact.url http://www.swagger.io/support
// @contact.email alao.adebusy@gmail.com, nabbo247@yahoo.com

// @license.name AddyTech Solution Ltd
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8060
// @BasePath /
// @schemes http
func main() {
	utilities.CreateLog()
	svc := gin.Default()

	url := ginSwagger.URL("http://localhost:8060/swagger/doc.json")

	svc.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	svc.GET("/", userRoute.CheckService)
	svc.POST("/api/users/createUser", userRoute.CreateNewUser)
	svc.POST("/api/users/Login", userRoute.Login)
	svc.GET("/api/users/GetAllUsers", userRoute.GetAllUsers)
	svc.GET("/api/users/GetUserDetailsByEmail/:email", userRoute.GetUserDetailsByEmail)

	svc.GET("/api/company/CheckEmailWithAuthCode/:email/:authcode", userRoute.CheckEmailWithAuthCode)
	svc.POST("/api/company/CompanyRegistration", userRoute.RegisterCompany)
	svc.GET("/api/company/GetCompanyDetail/:Email/:CompanyName", userRoute.GetCompanyDetail)
	svc.Run(":8060")
}
