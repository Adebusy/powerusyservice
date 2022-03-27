package controllersroute

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type shipperService struct {
	DbGorm *gorm.DB
}

func NewShippersRoute(db *gorm.DB) IShipperRoute {
	return &shipperService{db}
}

type IShipperRoute interface {
	UploadShippersDocument(ctx *gin.Context)
	ApproveShippersDocument(ctx *gin.Context)
	GetAllShippersDocument(ctx *gin.Context)
	GetShippersDocumentByCompanyname(ctx *gin.Context)
}

// UploadShippersDocument godoc
// @Summary Upload shippers document
// @Produce json
// @Tags shipper
// @Param shipperDoc body models.ShipperIn true "User Details"
// @Success 200 {object} models.ResponseMessage
// @Router /api/shippers/UploadShippersDocument [post]
func (ts shipperService) UploadShippersDocument(ctx *gin.Context) {
	requestBody := models.ShipperIn{}
	requestObj := models.Tbl_shipper{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Invalid requestbody uploaded, please re-confirm and try again later."
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	//cjeck if file already uploaded
	getshipper, err := newshipper.GetShippersDocument(requestBody.Companyname)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Error occurred. Please try again later."
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	if getshipper.Id != 0 && getshipper.Statusid != "0" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Shippers document had already been uploaded for this user."
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
	}

	//check company name
	checkCompName := newregistered.GetCompanyByCompanyName(strings.ToUpper(requestBody.Companyname))

	if checkCompName.Companyname == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "It appears this company  has not been created"
		//fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	//upload the document
	requestObj.Companyname = requestBody.Companyname
	requestObj.Userid = requestBody.Userid
	requestObj.Tin = requestBody.Tin
	requestObj.Tinpassword = requestBody.Tinpassword
	requestObj.Location = requestBody.Location
	requestObj.Phonenumber = requestBody.Phonenumber
	requestObj.Statusid = "0"

	uploadRequest, err := newshipper.UploadShippersDocument(requestObj)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Error occurred, please try again later " + err.Error()
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if uploadRequest != 0 {
		ResponBody.ResponseCode = "00"
		ResponBody.ResponseMessage = "Shippers document uploaded successfully"
		ctx.JSON(http.StatusOK, ResponBody)
	}
}

// ApproveShippersDocument godoc
// @Summary Approve shippers document
// @Produce json
// @Tags shipper
// @Param shipperDoc body models.ShipperApprovalReqIn true "User Details"
// @Success 200 {object} models.ResponseMessage
// @Router /api/shippers/ApproveShippersDocument [patch]
func (ts shipperService) ApproveShippersDocument(ctx *gin.Context) {
	requestBody := models.ShipperApprovalReqIn{}
	shipper := models.Tbl_shipper{}

	if err := ctx.ShouldBind(&requestBody); err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Error occurred, please try again later " + err.Error()
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	//check if initiator can approve request
	getService, err := newuserService.GetUserById(requestBody.Userid)

	if getService.Username == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "invalid approval ID supplied"
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	if getService.Roleid != 1 {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Approval does not have permission to perform this task"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	// check request
	uploadRequest, err := newshipper.GetShippersDocument(strings.ToUpper(requestBody.Companyname))
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Error occurred, please try again later " + err.Error()
		fmt.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if uploadRequest.Companyname == "" {
		ResponBody.ResponseCode = "00"
		ResponBody.ResponseMessage = "Shippers document has not been uploaded."
		ctx.JSON(http.StatusOK, ResponBody)
		return
	}
	//approve update
	shipper.Comment = requestBody.Comment
	shipper.Approvedby = requestBody.Approvedby
	shipper.Companyname = requestBody.Companyname
	shipper.Statusid = strconv.Itoa(requestBody.Statusid)
	updateQuery, err := newshipper.UpdateShipperDocument(&shipper)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		ctx.JSON(http.StatusOK, ResponBody)
		return
	}
	if updateQuery > 0 {
		ResponBody.ResponseCode = "00"
		ResponBody.ResponseMessage = "Shippers document updated successfully."
		ctx.JSON(http.StatusOK, ResponBody)
	}
}

// GetAllShippersDocument godoc
// @Summary Gets all shippers document
// @Produce json
// @Tags shipper
// @Success 200 {object} []models.ShipperOut
// @Router /api/shippers/GetAllShippersDocument [get]
func (ts shipperService) GetAllShippersDocument(ctx *gin.Context) {
	// check request
	uploadRequest, err := newshipper.GetAllShippersDocument()
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Error occurred, please try again later " + err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	ctx.JSON(http.StatusBadRequest, uploadRequest)
}

// GetShippersDocumentByCompanyname godoc
// @Summary Gets shipper's document by company name
// @Produce json
// @Tags shipper
// @Success 200 {object} []models.ShipperOut
// @Router /api/shippers/GetShippersDocumentByCompanyname/{comapanyname} [get]
func (ts shipperService) GetShippersDocumentByCompanyname(ctx *gin.Context) {
	compname := ctx.Param("companyname")
	uploadRequest, err := newshipper.GetShippersDocument(strings.ToUpper(compname))
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Error occurred, please try again later " + err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	if uploadRequest.Id == 0 {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Company documents have not been uploaded."
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	ctx.JSON(http.StatusOK, uploadRequest)
}
