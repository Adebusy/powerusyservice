package sqlserver

import (
	"strings"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/jinzhu/gorm"
)

func NewShipperService(db *gorm.DB) IShipper {
	return &dbconnect{db}
}

type IShipper interface {
	UploadShippersDocument(shipper models.Tbl_shipper) (int, error)
	GetShippersDocument(Companyname string) (models.ShipperOut, error)
	UpdateShipperDocument(requestBody *models.Tbl_shipper) (int, error)
	GetAllShippersDocument() ([]models.ShipperOut, error)
}

func (db dbconnect) UploadShippersDocument(shipper models.Tbl_shipper) (int, error) {
	err := db.DbGorm.Table(`Tbl_shipper`).Create(&shipper).Error
	if err != nil {
		return 0, err
	}
	return shipper.Id, nil
}

func (db dbconnect) GetShippersDocument(Companyname string) (models.ShipperOut, error) {
	retShipper := models.ShipperOut{}
	err := db.DbGorm.Table(`Tbl_shipper`).Where(`Companyname=?`, strings.ToUpper(Companyname)).First(&retShipper).Error
	if err != nil {
		if err.Error() != "record not found" {
			return retShipper, err
		}
	}
	return retShipper, nil
}

func (db dbconnect) GetAllShippersDocument() ([]models.ShipperOut, error) {
	retShipper := []models.ShipperOut{}
	err := db.DbGorm.Table(`Tbl_shipper`).First(&retShipper).Error
	if err != nil {
		if err.Error() != "record not found" {
			return retShipper, err
		}
	}
	return retShipper, nil
}

func (db dbconnect) UpdateShipperDocument(requestBody *models.Tbl_shipper) (int, error) {
	query := db.DbGorm.Table(`Tbl_shipper`).Where(`companyname=?`, requestBody.Companyname).UpdateColumns(map[string]interface{}{`statusid`: requestBody.Statusid, `approvedby`: requestBody.Approvedby, `comment`: requestBody.Comment})
	return int(query.RowsAffected), query.Error
}
