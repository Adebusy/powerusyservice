package controllersroute

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Adebusy/powerusyservice/models"
	util "github.com/Adebusy/powerusyservice/utilities"
	"github.com/mcnijman/go-emailaddress"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

// CheckService godoc
// @Summary Shows the new status of server.
// @Description Gets the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (ts userService) CheckService(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "I am up and runnings")
}

// CreateNewUser godoc
// @Summary Creates new user
// @Produce json
// @Tags user
// @Param user body models.UserIn true "User Details"
// @Success 200 {object} models.UserOut
// @Router /api/users/CreateNewUser [post]
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
	newusr.Roleid = userObj.Roleid
	newusr.Status = true
	newusr.Username = strings.ToUpper(userObj.Username)
	currentTimeUTC := time.Now().UTC()
	newusr.Dateadded = currentTimeUTC

	if _, err := emailaddress.Parse(newusr.Email); err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Invalid email address supplied."
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	retEmailDetail, _ := newuserService.GetUserByEmail(newusr.Email)
	if retEmailDetail.Id != 0 {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Email " + newusr.Email + "address already exists"
		ctx.JSON(http.StatusOK, ResponBody)
		return
	}

	retDetail := newuserService.GetUserDetailsByUsername(newusr.Username)
	if retDetail.Id != 0 {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Username already exist"
		ctx.JSON(http.StatusOK, ResponBody)
		return
	}

	createNewUser := newuserService.CreateUser(newusr)
	if createNewUser != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = createNewUser.Error()
		ctx.JSON(http.StatusOK, ResponBody)
		return
	}
	var useEmail []string
	useEmail = append(useEmail, newusr.Email)
	// send auth code to user
	sendAuth := util.SendEmail(ctx, useEmail, newusr.Authcode, ctx.Request)
	if sendAuth {
		ResponBody.ResponseCode = "00"
		ResponBody.ResponseMessage = fmt.Sprintf("User %s creaated successfully, authcode has been sent to the user email address", newusr.Username)
	} else {
		ResponBody.ResponseCode = "00"
		ResponBody.ResponseMessage = fmt.Sprintf("User %s creaated successfully, authcode will be sent shorltly to the user email address", newusr.Username)
	}
	ctx.JSON(http.StatusOK, ResponBody)
	return
}

// GetUserDetailsByEmail godoc
// @Summary gets user details by email address.
// @Description gets user details by email address.
// @Tags user
// @Accept */*
// @Produce json
// @Success 200 {object} models.Tbl_users
// @Router /api/users/GetUserDetailsByEmail/{email} [get]
func (ts userService) GetUserDetailsByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	if _, err := emailaddress.Parse(email); err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Invalid email address supplied."
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if _, err := emailaddress.Parse(email); err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Invalid email address supplied." + email
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	retQuery, _ := newuserService.GetUserByEmail(strings.ToUpper(email))
	// if retQuery.Error() != "" {
	// 	util.LogError(retQuery)
	// 	log.Fatal(retQuery.Error())
	// }
	ctx.JSON(http.StatusOK, retQuery)

}

// GetUserDetailsByUsername godoc
// @Summary Gets user details by username.
// @Description Gets user details by username.
// @Tags user
// @Accept */*
// @Produce json
// @Success 200 {object} models.Tbl_users
// @Router /api/users/GetUserDetailsByUsername/{username} [get]
func GetUserDetailsByUsername(ctx *gin.Context) models.Tbl_users {
	username := ctx.Query("username")
	retQuery := newuserService.GetUserDetailsByUsername(strings.ToUpper(username))
	// if retQuery.Error() != "" {
	// 	log.Fatal(retQuery.Error())
	// }
	return retQuery
}

// Login godoc
// @Summary log user into the syste
// @Produce json
// @Tags user
// @Param user body models.LoginIn true "Login user"
// @Success 200 {object} models.UserOut
// @Router /api/users/Login [post]
func (ts userService) Login(ctx *gin.Context) {
	requestBody := models.LoginIn{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		util.LogError(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	retQuery := newuserService.GetUserDetailsByUsernameAndPassword(strings.ToUpper(requestBody.Username), requestBody.Password)
	ctx.JSON(http.StatusOK, retQuery)
}

// GetAllUsers godoc
// @Summary gets list of all users
// @Produce json
// @Tags user
// @Success 200 {object} []models.UserOut
// @Router /api/users/GetAllUsers [get]
func (ts userService) GetAllUsers(ctx *gin.Context) {
	users := newuserService.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}

var ResponBody = models.ResponseMessage{}

// CheckEmailWithAuthCode godoc
// @Summary validates email with authentication code
// @Produce json
// @Tags user
// @Success 200 {object} models.ResponseMessage
// @Router /api/users/CheckEmailWithAuthCode/{email}/{authcode}  [get]
func (ts userService) CheckEmailWithAuthCode(ctx *gin.Context) {
	reponseUser := models.Tbl_users{}
	email := ctx.Param("email")
	authcode := ctx.Param("authcode")
	if _, err := emailaddress.Parse(email); err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Invalid email address supplied."
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if email == "" || authcode == "" {
		ctx.JSON(http.StatusBadRequest, "Parameter Email and password is required")
	}

	retQuery := ts.DbGorm.Debug().Table("Tbl_users").Where("authcode=? and email=?", authcode, strings.ToUpper(email)).Find(&reponseUser).Error
	if retQuery != nil {
		if retQuery.Error() != "" && retQuery.Error() == "record not found" {
			ResponBody.ResponseCode = "01"
			ResponBody.ResponseMessage = "Invalid email or authentication code supplied."
			fmt.Println(retQuery.Error())
			ctx.JSON(http.StatusBadRequest, ResponBody)
			return
		}

		if retQuery.Error() != "" && retQuery.Error() != "record not found" {
			ResponBody.ResponseCode = "01"
			ResponBody.ResponseMessage = retQuery.Error()
			fmt.Println(retQuery.Error())
			ctx.JSON(http.StatusBadRequest, ResponBody)
			return
		}
	}

	respoReq := newuserService.CheckEmailWithAuthCode(strings.ToUpper(email), strings.ToUpper(authcode))

	if respoReq.Username != "" {
		//change profile status
		if newuserService.UpdateUserStatusUponAuthCode(strings.ToUpper(email), strings.ToUpper(authcode), 2) {
			ResponBody.ResponseCode = "00"
			ResponBody.ResponseMessage = "Authentication was successful"
			ctx.JSON(http.StatusOK, ResponBody)
			return
		}
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Authentication was not successful, kindly try again later."
		ctx.JSON(http.StatusBadRequest, ResponBody)
	}
}
