package sqlserver

import (
	"fmt"
	"time"

	"github.com/Adebusy/powerusyservice/app/driver"
	"github.com/Adebusy/powerusyservice/models"
	"github.com/jinzhu/gorm"
)

type IKYC interface {
	GetKYCByCompanyId(companyId int) (models.KYCOut, error)
	CreateKYC(requestObj models.Tbl_KYC) (int, error)
	GetAllKYC() ([]models.KYCsOut, error)
	ApproveKYC(KYCApprovalIn models.KYCApprovalIn) (int, error)
}

func NewKYCService(db *gorm.DB) IKYC {
	return &dbconnect{db}
}

var newregistered = NewRegistered(driver.GetDB())

func (db dbconnect) GetKYCByCompanyId(companyId int) (models.KYCOut, error) {
	retVal := models.KYCOut{}
	err := db.DbGorm.Table(`tbl_kyc`).Where(`registeredid=?`, companyId).First(&retVal).Error
	if err != nil {
		fmt.Println(err.Error())
		return retVal, err
	}
	return retVal, nil
}

func (db dbconnect) CreateKYC(requestObj models.Tbl_KYC) (int, error) {
	doinser := db.DbGorm.Table(`tbl_kyc`).Create(&requestObj).Error
	if doinser != nil {
		return 0, doinser
	}
	return requestObj.Id, nil
}

func (db dbconnect) GetAllKYC() ([]models.KYCsOut, error) {
	retVal := []models.KYCsOut{}
	err := db.DbGorm.Debug().Table(`tbl_kyc`).Find(&retVal).Error
	if err != nil {
		fmt.Println(err.Error())
		return retVal, err
	}
	if len(retVal) != 0 {
		for i, _ := range retVal {
			retCompant, _ := newregistered.GetCompanyByCompanyId(retVal[i].Registeredid)
			retVal[i].Companyname = retCompant.Companyname
		}
	}
	return retVal, nil
}

func (db dbconnect) ApproveKYC(KYCApprovalIn models.KYCApprovalIn) (int, error) {
	doinser := db.DbGorm.Debug().Table(`tbl_kyc`).Where(`registeredid=?`, KYCApprovalIn.Registeredid).Updates(map[string]interface{}{"Approvalcomment": &KYCApprovalIn.Approvalcomment, "statusid": &KYCApprovalIn.Status, "Approvedby": &KYCApprovalIn.Approvedby, "Dateapproved": time.Now()}).Error
	if doinser != nil {
		return 0, doinser
	}
	return 1, nil
}
