package controllersroute

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/swaggo/swag/example/celler/httputil"
)

type userService struct {
	DbGorm *gorm.DB
}

func NewIuserController(db *gorm.DB) IuserController {
	return &userService{db}
}

type IuserController interface {
	CreateNewUser(ctx *gin.Context)
	CheckService(ctx *gin.Context)
	LoginUSer(ctx *gin.Context)
}

func (ts userService) CheckService(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "I am up and runnings")
}

func (ts userService) CreateNewUser(ctx *gin.Context) {
	userObj := &models.UserIn{}
	respon := models.ResponseMessage{}
	newusr := &models.Tbl_users{}

	if err := ctx.ShouldBindJSON(userObj); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newusr.Authcode = "authcode"
	newusr.Email = strings.ToUpper(userObj.Email)
	newusr.Firstname = strings.ToUpper(userObj.FirstName)
	newusr.Lastname = strings.ToUpper(userObj.LastName)
	newusr.Middlename = strings.ToUpper(userObj.MiddleName)
	newusr.Password = userObj.Password
	newusr.Phonenumber = userObj.PhoneNumber
	newusr.Roleid = 1
	newusr.Status = true
	newusr.Username = strings.ToUpper(userObj.Username)
	currentTimeUTC := time.Now().UTC()
	newusr.Dateadded = currentTimeUTC

	retEmailDetail := ts.GetUserDetailsByEmail(newusr.Email)
	if len(retEmailDetail.Email) != 0 {
		respon.ResponseCode = 01
		respon.ResponseMessage = "Email " + retEmailDetail.Email + "address already exists"
		ctx.JSON(http.StatusOK, respon)
		return
	}

	retDetail := ts.GetUserDetailsByUsername(newusr.Username)
	if len(retDetail.Firstname) != 0 {
		respon.ResponseCode = 01
		respon.ResponseMessage = "Username already exist"
		ctx.JSON(http.StatusOK, respon)
		return
	}

	inst := ts.DbGorm.Debug().Table("Tbl_users").Create(&newusr).Error
	if inst != nil {
		respon.ResponseCode = 01
		respon.ResponseMessage = inst.Error()
		ctx.JSON(http.StatusOK, respon)
		return
	}
	respon.ResponseCode = 00
	respon.ResponseMessage = fmt.Sprintf("User %s creaated successfully", newusr.Username)
	ctx.JSON(http.StatusOK, respon)
	return
}

func (ts userService) GetUserDetailsByEmail(emailAddress string) models.Tbl_users {
	fmt.Println("get here GetUserDetailsByEmail")
	fmt.Println(emailAddress)
	newusr := models.Tbl_users{}
	ts.DbGorm.Table(`Tbl_users`).Where(`Email =?`, emailAddress).First(&newusr)
	return newusr
}

func (ts userService) GetUserDetailsByUsername(username string) models.Tbl_users {
	reponseUser := models.Tbl_users{}
	ts.DbGorm.Table(`Tbl_users`).Where(`username =?`, username).First(&reponseUser)
	return reponseUser
}

func (ts userService) LoginUSer(ctx *gin.Context) {
	requestBody := models.LoginIn{}
	responseBody := models.UserOut{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ts.DbGorm.Debug().Table(`Tbl_users`).Where(`username=? and password=?`, requestBody.Username, requestBody.Password).First(&responseBody)
	ctx.JSON(http.StatusOK, responseBody)
}
