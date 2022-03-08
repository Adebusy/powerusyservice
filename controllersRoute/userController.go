package controllersroute

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Adebusy/powerusyservice/models"
	util "github.com/Adebusy/powerusyservice/utilities"

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
	GetAllUser(ctx *gin.Context)
	CheckEmailWithAuthCode(ctx *gin.Context)
}

func (ts userService) CheckService(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "I am up and runnings")
}

func (ts userService) CreateNewUser(ctx *gin.Context) {
	userObj := &models.UserIn{}
	newusr := &models.Tbl_users{}

	if err := ctx.ShouldBindJSON(userObj); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newusr.Authcode = util.GenerateAuthCode()
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
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Email " + retEmailDetail.Email + "address already exists"
		ctx.JSON(http.StatusOK, ResponBody)
		return
	}

	retDetail := ts.GetUserDetailsByUsername(newusr.Username)
	if len(retDetail.Firstname) != 0 {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Username already exist"
		ctx.JSON(http.StatusOK, ResponBody)
		return
	}

	inst := ts.DbGorm.Table("Tbl_users").Create(&newusr).Error
	if inst != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = inst.Error()
		ctx.JSON(http.StatusOK, ResponBody)
		return
	}
	ResponBody.ResponseCode = "00"
	ResponBody.ResponseMessage = fmt.Sprintf("User %s creaated successfully", newusr.Username)
	ctx.JSON(http.StatusOK, ResponBody)
	return
}

func (ts userService) GetUserDetailsByEmail(emailAddress string) models.Tbl_users {
	newusr := models.Tbl_users{}
	retQuery := ts.DbGorm.Table(`Tbl_users`).Where(`Email =?`, emailAddress).First(&newusr).Error
	if retQuery.Error() != "" {
		log.Fatal(retQuery.Error())
	}

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
	ts.DbGorm.Table(`Tbl_users`).Where(`username=? and password=?`, requestBody.Username, requestBody.Password).First(&responseBody)
	ctx.JSON(http.StatusOK, responseBody)
}
func (ts userService) GetAllUser(ctx *gin.Context) {
	users := []models.UserOut{}
	ts.DbGorm.Table(`Tbl_users`).Find(&users)
	ctx.JSON(http.StatusOK, users)
}

var ResponBody = models.ResponseMessage{}

func (ts userService) CheckEmailWithAuthCode(ctx *gin.Context) {
	ret := ctx.Request.URL
	fmt.Println(ret)
	reponseUser := models.Tbl_users{}
	email := ctx.Query("email")
	authcode := ctx.Query("authcode")

	if email == "" || authcode == "" {
		ctx.JSON(http.StatusBadRequest, "Parameter Email and password is required")
	}
	ts.DbGorm.Table("Tbl_users").Where("authcode=? and email=?", authcode, strings.ToUpper(email)).Find(&reponseUser)
	if reponseUser.Username != "" {
		ResponBody.ResponseCode = "00"
		ResponBody.ResponseMessage = "Authentication was successful"
		ctx.JSON(http.StatusOK, ResponBody)
	}
}
