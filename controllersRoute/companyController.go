// package controllersroute this is used to manage all company activities
package controllersroute

import (
	"fmt"
	"net/http"
	"strconv"
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

var (
	newuserService = sqlserver.NewUser(driver.GetDB())
	newregistered  = sqlserver.NewRegistered(driver.GetDB())
	newKYC         = sqlserver.NewKYCService(driver.GetDB())
	newshipper     = sqlserver.NewShipperService(driver.GetDB())
)

type IuserController interface {
	CreateNewUser(ctx *gin.Context)
	CheckService(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	GetUserDetailsByEmail(ctx *gin.Context)
	CheckEmailWithAuthCode(ctx *gin.Context)

	RegisterCompany(ctx *gin.Context)
	GetCompanyDetailByEmailandCompName(ctx *gin.Context)
	UploadCompanyDocuments(ctx *gin.Context)
	GetCompanyByCompanyName(ctx *gin.Context)

	UploadKYC(ctx *gin.Context)
	GetKYCbyCompanyId(ctx *gin.Context)
	GetAllKYC(ctx *gin.Context)
	ApproveKYC(ctx *gin.Context)
}

// RegisterCompany godoc
// @Summary Register new company
// @Produce json
// @Tags company
// @Param company body models.CompanyDetailsIn true "Company details"
// @Success 200 {object} models.ResponseMessage
// @Router /api/company/RegisterCompany [post]
func (ts userService) RegisterCompany(ctx *gin.Context) {
	requestBody := models.CompanyDetailsIn{}
	Registered := models.Tbl_Registered{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		//util.LogError(err)
		//log.Panic(err)
		ctx.JSON(http.StatusBadRequest, err)
	}

	docheck := newregistered.GetCompanyByCompanyName("testing")

	fmt.Println(docheck)

	//check profile exist
	retUser, err := newuserService.GetUserByEmailandUserId(strings.ToUpper(requestBody.Email), requestBody.UserId)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		//util.LogError(err)
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	fmt.Println(retUser.Username)
	if retUser.Username == "" {
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
	//register company details
	createComp, err := newregistered.RegisterCompany(Registered)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		//util.LogError(query)
		return
	}
	//upload company detail

	ResponBody.ResponseCode = "00"
	ResponBody.ResponseMessage = fmt.Sprintf("Document uploaded successfully with record id %d", createComp.Id)
	ctx.JSON(http.StatusBadRequest, ResponBody)
}

// GetCompanyDetailByEmailandCompName godoc
// @Summary Gets company details by email address and company name.
// @Produce json
// @Tags company
// @Success 200 {object} models.CompanyDetailsOut
// @Router /api/company/GetCompanyDetailByEmailandCompName/{email}/{companyname} [get]
func (ts userService) GetCompanyDetailByEmailandCompName(ctx *gin.Context) {
	email := ctx.Param("email")
	CompanyName := ctx.Param("companyname")
	CompanyDocumentOut := models.CompanyDetailsOut{}
	reponseUser := models.Tbl_users{}

	if err := ts.DbGorm.Table("Tbl_Users").Where("Email=? ", strings.ToUpper(email)).Find(&reponseUser).Error; err != nil {
		if err.Error() != "" && err.Error() != "record not found" {
			ResponBody.ResponseCode = "01"
			ResponBody.ResponseMessage = err.Error()
			//util.LogError(err)
			ctx.JSON(http.StatusBadRequest, ResponBody)
			return
		}
	}

	if reponseUser.Username == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Invalid user details supplied"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	query := ts.DbGorm.Table("Tbl_Registered").Where("CompanyName=?", strings.ToUpper(CompanyName)).Find(&CompanyDocumentOut).Error
	if query != nil {
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
	ctx.JSON(http.StatusBadRequest, CompanyDocumentOut)
}

// UploadCompanyDocuments godoc
// @Summary Upload company document
// @Produce json
// @Tags company
// @Param CompanyDoc body models.ImportationDocumentIn true "Upload company document"
// @Success 200 {object} models.ResponseMessage
// @Router /api/company/UploadCompanyDocuments [post]
func (ts userService) UploadCompanyDocuments(ctx *gin.Context) {
	req := models.ImportationDocumentIn{}
	ctx.ShouldBindJSON(&req)
	if req.Documentname == "" || req.Documentpath == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Document name is required."
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	//check if company had been created
	compId, err := strconv.Atoi(req.Importationid)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "CompanyId must be integer"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	rec, err := newregistered.GetCompanyByCompanyId(compId)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	fmt.Println(rec.Accountnumber)
	if rec.Companyname == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "invalid company id supplied"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	docuRec := models.Tbl_ImportationDocument{
		Importationid: req.Importationid,
		Documentname:  req.Documentname,
		Documentpath:  req.Documentpath,
		Dateadded:     time.Now(),
		Statusid:      0,
	}

	//do upload
	ret, err := newregistered.UploadDocument(docuRec)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	if ret == 1 && err == nil {
		ResponBody.ResponseCode = "00"
		ResponBody.ResponseMessage = "Document uploaded successfully."
		ctx.JSON(http.StatusOK, ResponBody)
	}
}

// GetCompanyByCompanyName godoc
// @Summary Gets company details by company name.
// @Produce json
// @Tags company
// @Success 200 {object} models.CompanyDetailsOut
// @Router /api/company/GetCompanyByCompanyName/{CompanyName} [get]
func (ts userService) GetCompanyByCompanyName(ctx *gin.Context) {
	CompanyName := ctx.Param("companyname")
	getCompany := newregistered.GetCompanyByCompanyName(CompanyName)
	ctx.JSON(http.StatusOK, getCompany)
}

// UploadKYC godoc
// @Summary Upload KYC document.
// @Produce json
// @Tags company
// @Param KYC body models.KYCIn true "KYC document"
// @Success 200 {object} models.ResponseMessage
// @Router /api/company/UploadKYC [post]
func (ts userService) UploadKYC(ctx *gin.Context) {
	requestBody := models.KYCIn{}
	requestObj := models.Tbl_KYC{}
	ctx.ShouldBindJSON(&requestBody)

	rec, err := newregistered.GetCompanyByCompanyId(requestBody.Registeredid)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if rec.Companyname == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "invalid company id supplied"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if requestBody.Articlesofassociation == "" || requestBody.Auditedfinancialstatement == "" || requestBody.Certificateofincorporation == "" || requestBody.Memorandomofassociation == "" || requestBody.Powerofattorneygranted == "" {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "Audited financial statement, Certificate of incorporation, Memorandom of association, and Power of attorney granted are required for KYC to be performed."
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	//check file already been uploaded
	retVal, err := newKYC.GetKYCByCompanyId(requestBody.Registeredid)
	if err != nil {
		if err.Error() != "record not found" {
			ResponBody.ResponseCode = "01"
			ResponBody.ResponseMessage = err.Error()
			ctx.JSON(http.StatusBadRequest, ResponBody)
			return
		}
	}

	if retVal.Id != 0 {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "KYC documents had already been uploaded for this user"
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	//insert
	requestObj.Approvalcomment = "submitted"
	requestObj.Articlesofassociation = requestBody.Articlesofassociation
	requestObj.Auditedfinancialstatement = requestBody.Auditedfinancialstatement
	requestObj.Certificateofincorporation = requestBody.Certificateofincorporation
	requestObj.Dateadded = time.Now()
	requestObj.Memorandomofassociation = requestBody.Memorandomofassociation
	requestObj.Powerofattorneygranted = requestBody.Powerofattorneygranted
	requestObj.Registeredid = requestBody.Registeredid
	requestObj.Statusid = 0
	requestObj.Taxclearancecertificate = requestBody.Taxclearancecertificate
	requestObj.Validbusinesslicense = requestBody.Validbusinesslicense
	requestObj.Dateapproved = time.Now()

	//do insert
	queryId, err := newKYC.CreateKYC(requestObj)
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = "error occurred, please try again later " + err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}

	if queryId != 0 {
		ResponBody.ResponseCode = "00"
		ResponBody.ResponseMessage = "KYC document uploaded successfully."
		ctx.JSON(http.StatusOK, ResponBody)
	}
}

// GetKYCbyCompanyId godoc
// @Summary Gets company KYC details by company Id.
// @Produce json
// @Tags company
// @Success 200 {object} models.KYCOut
// @Router /api/company/GetKYCbyCompanyId/{Id} [get]
func (ts userService) GetKYCbyCompanyId(ctx *gin.Context) {
	requuestID, _ := strconv.Atoi(ctx.Param("Id"))
	if requuestID != 0 {
		querykyc, err := newKYC.GetKYCByCompanyId(requuestID)
		if err != nil {
			ResponBody.ResponseCode = "01"
			ResponBody.ResponseMessage = err.Error()
			ctx.JSON(http.StatusBadRequest, ResponBody)
			return
		}
		ctx.JSON(http.StatusOK, querykyc)
	}
}

// GetAllKYC godoc
// @Summary Gets all companies KYC details.
// @Produce json
// @Tags company
// @Success 200 {object} []models.KYCsOut
// @Router /api/company/GetAllKYC [get]
func (ts userService) GetAllKYC(ctx *gin.Context) {
	querykyc, err := newKYC.GetAllKYC()
	if err != nil {
		ResponBody.ResponseCode = "01"
		ResponBody.ResponseMessage = err.Error()
		ctx.JSON(http.StatusBadRequest, ResponBody)
		return
	}
	ctx.JSON(http.StatusOK, querykyc)
}

// ApproveKYC godoc
// @Summary Approve KYC document.
// @Produce json
// @Tags company
// @Param KYC body models.KYCApprovalIn true "KYC approval"
// @Success 200 {object} models.ResponseMessage
// @Router /api/company/ApproveKYC [post]
func (ts userService) ApproveKYC(ctx *gin.Context) {
	fmt.Println("get here noew")
	requestObj := models.KYCApprovalIn{}
	err := ctx.ShouldBindJSON(&requestObj)
	if err == nil {
		querykyc, err := newKYC.ApproveKYC(requestObj)
		if err != nil {
			ResponBody.ResponseCode = "01"
			ResponBody.ResponseMessage = err.Error()
			ctx.JSON(http.StatusBadRequest, ResponBody)

		} else {
			if querykyc != 0 {
				ResponBody.ResponseCode = "00"
				ResponBody.ResponseMessage = "KYC approved successfully"
				ctx.JSON(http.StatusBadRequest, ResponBody)
				return
			}
		}
	}
	ResponBody.ResponseCode = "01"
	ResponBody.ResponseMessage = "error occurred, please try again later " + err.Error()
	ctx.JSON(http.StatusBadRequest, ResponBody)
}
