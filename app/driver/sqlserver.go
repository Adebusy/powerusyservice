package driver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/Adebusy/powerusyservice/utilities"
	ut "github.com/Adebusy/powerusyservice/utilities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
)

var DbGorm *gorm.DB
var ErrGorm error

func init() {
	godotenv.Load()
	var dbStatus models.ConfigStruct
	//connectionString := fmt.Sprintf("sqlserver://%s:%s@localhost:1433?database=%s", utilities.GoDotEnvVariable("UserID"), utilities.GoDotEnvVariable("Password"), utilities.GoDotEnvVariable("Database"))
	connectionString := fmt.Sprintf("sqlserver://%s:%s@plesk2500.is.cc?database=%s", utilities.GoDotEnvVariable("UserID"), utilities.GoDotEnvVariable("Password"), utilities.GoDotEnvVariable("Database"))
	fmt.Println(connectionString)
	DbGorm, ErrGorm = gorm.Open("mssql", connectionString)
	if ErrGorm != nil {
		fmt.Println(ErrGorm.Error())
		ut.LogError(ErrGorm)
	}

	read, err := ioutil.ReadFile("config.json")
	if err != nil {
		ut.LogError(err)
	}
	if err := json.Unmarshal(read, &dbStatus); err != nil {
		ut.LogError(err)
	}
	fmt.Println(dbStatus.IsDropExistingTables)
	if dbStatus.IsDropExistingTables {
		DbGorm.Debug().DropTableIfExists()
		DbGorm.Debug().DropTableIfExists(&models.Tbl_users{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_Registered{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_ImportationDocument{})

		DbGorm.Debug().DropTableIfExists(&models.Tbl_KYC{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_role{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_status{})

		DbGorm.Debug().DropTableIfExists(&models.Tbl_auditlog{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_clearing{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_goodstype{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_importation{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_servicetype{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_shipper{})
		DbGorm.Debug().DropTableIfExists(&models.Tbl_subgoodstype{})
	}

	if dbStatus.CreateTable {
		DbGorm.SingularTable(true)
		DbGorm.Debug().AutoMigrate(&models.Tbl_users{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_Registered{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_ImportationDocument{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_KYC{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_role{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_status{})

		DbGorm.Debug().AutoMigrate(&models.Tbl_auditlog{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_clearing{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_goodstype{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_importation{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_servicetype{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_shipper{})
		DbGorm.Debug().AutoMigrate(&models.Tbl_subgoodstype{})
	}

	dbStatus.IsDropExistingTables = false
	dbStatus.CreateTable = false
	domarchal, _ := json.Marshal(dbStatus)
	_ = ioutil.WriteFile("config.json", domarchal, 0400)
}

func GetDB() *gorm.DB {
	return DbGorm
}
