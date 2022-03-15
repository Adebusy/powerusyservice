package sqlserver

import (
	"fmt"
	"strings"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/jinzhu/gorm"
)

type dbconnect struct {
	DbGorm *gorm.DB
}

func NewRegistered(db *gorm.DB) IRegistered {
	return &dbconnect{db}
}

type IRegistered interface {
	GetCompanyByCompanyName(CompanyName string) models.CompanyDocumentOut
	GetCompanyByCompanyNameAndUserId(UserId, StatusId int, CompanyName string) (models.CompanyDocumentOut, error)
	RegisterCompany(register models.Tbl_Registered) (models.Tbl_Registered, error)
}

func (db dbconnect) GetCompanyByCompanyName(CompanyName string) models.CompanyDocumentOut {
	company := models.CompanyDocumentOut{}
	if retQuery := db.DbGorm.Debug().Table("Tbl_Registered").Where("CompanyName=?", strings.ToUpper(CompanyName)).Find(&company).Error; retQuery != nil {
		fmt.Println(retQuery)
	}
	return company
}

func (db dbconnect) GetCompanyByCompanyNameAndUserId(UserId, StatusId int, CompanyName string) (models.CompanyDocumentOut, error) {
	company := models.CompanyDocumentOut{}
	queryCheck := db.DbGorm.Debug().Table("Tbl_Registered").Where("UserId =? and CompanyName=? and StatusId=?", UserId, CompanyName, StatusId).Select(&company).Error
	return company, queryCheck
}

func (db dbconnect) RegisterCompany(register models.Tbl_Registered) (models.Tbl_Registered, error) {
	retQuery := db.DbGorm.Debug().Table("Tbl_Registered").Create(&register).Error
	return register, retQuery
}
