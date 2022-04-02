package sqlserver

import (
	"fmt"
	"strings"

	"github.com/Adebusy/powerusyservice/models"
	ut "github.com/Adebusy/powerusyservice/utilities"
	"github.com/jinzhu/gorm"
)

type dbconnect struct {
	DbGorm *gorm.DB
}

func NewRegistered(db *gorm.DB) IRegistered {
	return &dbconnect{db}
}

type IRegistered interface {
	GetCompanyByCompanyName(CompanyName string) models.CompanyDetailsOut
	GetCompanyByCompanyNameAndUserId(UserId, StatusId int, CompanyName string) (models.CompanyDetailsOut, error)
	RegisterCompany(register models.Tbl_Registered) (models.Tbl_Registered, error)
	GetCompanyByCompanyId(UserId int) (models.CompanyDetailsOut, error)
	UploadDocument(compDocument models.Tbl_ImportationDocument) (int, error)
}

func (db dbconnect) GetCompanyByCompanyName(CompanyName string) models.CompanyDetailsOut {
	registered := models.CompanyDetailsOut{}
	if retQuery := db.DbGorm.Debug().Table(`Tbl_Registered`).Where(`CompanyName=?`, strings.ToUpper(CompanyName)).Find(&registered).Error; retQuery != nil {
		ut.LogError(retQuery)
	}
	return registered
}

func (db dbconnect) GetCompanyByCompanyNameAndUserId(UserId, StatusId int, CompanyName string) (models.CompanyDetailsOut, error) {
	company := models.CompanyDetailsOut{}
	queryCheck := db.DbGorm.Table(`Tbl_Registered`).Where(`UserId =? and CompanyName=? and StatusId=?`, UserId, CompanyName, StatusId).Select(&company).Error
	return company, queryCheck
}

func (db dbconnect) RegisterCompany(register models.Tbl_Registered) (models.Tbl_Registered, error) {
	retQuery := db.DbGorm.Table(`Tbl_Registered`).Create(&register).Error
	return register, retQuery
}

func (db dbconnect) GetCompanyByCompanyId(UserId int) (models.CompanyDetailsOut, error) {
	company := models.CompanyDetailsOut{}
	fmt.Println("get here 1")
	fmt.Println(UserId)
	queryCheck := db.DbGorm.Debug().Table(`Tbl_Registered`).Where(`Id =?`, UserId).First(&company).Error
	if queryCheck != nil {
		ut.LogError(queryCheck)
	}
	fmt.Println(company.Serviceid)
	return company, nil
}

func (db dbconnect) UploadDocument(compDocument models.Tbl_ImportationDocument) (int, error) {
	retQuery := db.DbGorm.Debug().Table(`tbl_Importation_Document`).Create(&compDocument).Error
	return compDocument.Id, retQuery
}
