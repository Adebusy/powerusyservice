// package controllersroute this is used to manage all company activities
package controllersroute

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Adebusy/powerusyservice/app/driver"
	"github.com/Adebusy/powerusyservice/datasource/sqlserver"
	"github.com/Adebusy/powerusyservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type userService struct {
	DbGorm *gorm.DB
}

func NewIuserController(db *gorm.DB) IuserController {
	return &userService{db}
}

var newuserService = sqlserver.NewUser(driver.GetDB())
var newregistered = sqlserver.NewRegistered(driver.GetDB())

type IuserController interface {
	CreateNewUser(ctx *gin.Context)
	CheckService(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	GetUserDetailsByEmail(ctx *gin.Context)
	CheckEmailWithAuthCode(ctx *gin.Context)

	RegisterCompany(ctx *gin.Context)
	GetCompanyDetail(ctx *gin.Context)
}

// RegisterCompany godoc
// @Summary Register new company
// @Produce json
// @Tags company
// @Param user body models.CompanyDocumentIn true "Company details"
// @Success 200 {object} models.ResponseMessage
// @Router /api/company/RegisterCompany [post]
func (ts userService) RegisterCompany(ctx *gin.Context) {
	requestBody := models.CompanyDocumentIn{}
	//CompanyDocumentOut := models.CompanyDocumentOut{}
	reponseUser := models.Tbl_users{}
	Registered := models.Tbl_Registered{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		//util.LogError(err)
		//log.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	// // validate company username and email address exist
	// if err := ts.DbGorm.Table("Tbl_Users").Where("Email=? and Id=?", requestBody.Email, requestBody.UserId).Find(&reponseUser).Error; err != nil {
	// 	if err.Error() != "" {
	// 		ResponBody.ResponseCode = "01"
	// 		ResponBody.ResponseMessage = err.Error()
	// 		//util.LogError(err)
	// 		ctx.JSON(http.StatusBadRequest, ResponBody)
	// 		return
	// 	}
	// }

	_, err := newuserService.GetUserByEmailandUserId(strings.ToUpper(requestBody.Email), requestBody.UserId)
	if err.Error() != "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		//util.LogError(err)
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if reponseUser.Username == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Invalid user details supplied"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	//insert company record
	Registered.Accountnumber = requestBody.AccountNumber
	Registered.Approvalcomment = "NEW REQUEST"
	Registered.Approvedby = 0
	Registered.Bankname = strings.ToUpper(requestBody.BankName)
	Registered.Companyaddress = strings.ToUpper(requestBody.CompanyAddress)
	Registered.Companylocation = strings.ToUpper(requestBody.CompanyLocation)
	Registered.Companylogo = requestBody.CompanyLogo
	Registered.Companyname = strings.ToUpper(requestBody.CompanyName)
	Registered.Description = strings.ToUpper(requestBody.Description)
	Registered.Postaddress = strings.ToUpper(requestBody.PostAddress)
	Registered.Serviceid = requestBody.ServiceId
	Registered.Statusid = 0
	Registered.Userid = requestBody.UserId
	Registered.Workingdays = requestBody.WorkingDays
	Registered.Workinghours = requestBody.WorkingHours
	Registered.Dateadded = time.Now()
	Registered.Dateapproved = time.Now()

	query, err := newregistered.GetCompanyByCompanyNameAndUserId(Registered.Userid, 0, Registered.Companyname)
	//queryCheck checks if document already been uploaded
	//queryCheck := ts.DbGorm.Debug().Table("Tbl_Registered").Where("UserId =? and CompanyName=? and StatusId=?", Registered.Userid, Registered.Companyname, 0).Select(&CompanyDocumentOut).Error
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		//util.LogError(queryCheck)
		return
	}

	if query.Accountnumber != "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Company documents had already been uploaded"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	//upload document
	//query := ts.DbGorm.Debug().Table("Tbl_Registered").Create(&Registered).Error
	createComp, err := newregistered.RegisterCompany(Registered)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		//util.LogError(query)
		return
	}
	ResponBody.ResponseCode = "00"
	ResponBody.ResponseMessage = fmt.Sprintf("Document uploaded successfully with record id %d", createComp.Id)
	ctx.JSON(http.StatusBadRequest, ResponBody)
}

// GetCompanyDetail godoc
// @Summary Gets company details by email address and company name.
// @Produce json
// @Tags company
// @Success 200 {object} models.CompanyDocumentOut
// @Router /api/company/GetComplaintByRefID/{Email}/{CompanyName} [get]
func (ts userService) GetCompanyDetail(ctx *gin.Context) {
	email := ctx.Param("Email")
	CompanyName := ctx.Param("CompanyName")
	CompanyDocumentOut := models.CompanyDocumentOut{}
	reponseUser := models.Tbl_users{}

	if err := ts.DbGorm.Debug().Table("Tbl_Users").Where("Email=? ", strings.ToUpper(email)).Find(&reponseUser).Error; err != nil {
		if err.Error() != "" && err.Error() != "record not found" {
			fmt.Println("record 1")
			ResponBody.ResponseCode = "01"
			ResponBody.ResponseMessage = err.Error()
			fmt.Println("record 2")
			//util.LogError(err)
			ctx.JSON(http.StatusBadRequest, ResponBody)
			return
		}
	}

	if reponseUser.Username == "" {
		fmt.Println("GetCompanyDetail4")
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Invalid user details supplied"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	fmt.Println("GetCompanyDetail5")
	query := ts.DbGorm.Debug().Table("Tbl_Registered").Where("CompanyName=?", strings.ToUpper(CompanyName)).Find(&CompanyDocumentOut).Error
	if query != nil {
		fmt.Println("GetCompanyDetail6")
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = query.Error()
		//util.LogError(err)
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if CompanyDocumentOut.Accountnumber == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Company documents have not been uploaded."
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	fmt.Println("GetCompanyDetail7")
	ctx.JSON(http.StatusBadRequest, CompanyDocumentOut)
}
