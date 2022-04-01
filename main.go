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
var shipperRoute = controllersroute.NewShippersRoute(driver.GetDB())

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
	//update test
	svc.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	svc.GET("/", userRoute.CheckService)                                                            //done
	svc.POST("/api/users/createUser", userRoute.CreateNewUser)                                      //done
	svc.POST("/api/users/Login", userRoute.Login)                                                   //done
	svc.GET("/api/users/GetAllUsers", userRoute.GetAllUsers)                                        //done
	svc.GET("/api/users/GetUserDetailsByEmail/:email", userRoute.GetUserDetailsByEmail)             //done
	svc.GET("/api/users/CheckEmailWithAuthCode/:email/:authcode", userRoute.CheckEmailWithAuthCode) //done

	svc.POST("/api/company/CompanyRegistration", userRoute.RegisterCompany)                                                      //done
	svc.GET("/api/company/GetCompanyDetailByEmailandCompName/:email/:companyname", userRoute.GetCompanyDetailByEmailandCompName) //done
	svc.POST("/api/company/UploadCompanyDocuments", userRoute.UploadCompanyDocuments)                                            //done
	svc.GET("/api/company/GetCompanyByCompanyName/:companyname", userRoute.GetCompanyByCompanyName)                              //done
	svc.POST("/api/company/UploadKYC", userRoute.UploadKYC)                                                                      //done
	svc.GET("/api/company/GetKYCbyCompanyId/:Id", userRoute.GetKYCbyCompanyId)                                                   //done
	svc.GET("/api/company/GetAllKYC", userRoute.GetAllKYC)                                                                       //done
	svc.POST("/api/company/ApproveKYC", userRoute.ApproveKYC)                                                                    //done

	svc.POST("/api/shippers/UploadShippersDocument", shipperRoute.UploadShippersDocument)                                 //done
	svc.PATCH("/api/shippers/ApproveShippersDocument", shipperRoute.ApproveShippersDocument)                              //done
	svc.GET("/api/shippers/GetAllShippersDocument", shipperRoute.GetAllShippersDocument)                                  //done
	svc.GET("/api/shippers/GetShippersDocumentByCompanyname/:companyname", shipperRoute.GetShippersDocumentByCompanyname) //done
	svc.Run(":8060")

	//runCommand := exec.Command(`curl -s https://api.github.com/repos/progrium/go-basher | json-pointer /owner/login`, `cp ./* ./testfile`)

	//runCommand.Run()
}
