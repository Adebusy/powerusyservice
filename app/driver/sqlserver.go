package driver

import (
	"fmt"
	"os"

	"github.com/Adebusy/powerusyservice/utilities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/joho/godotenv"
)

var DbGorm *gorm.DB
var ErrGorm error

func init() {
	godotenv.Load()
	connectionString := fmt.Sprintf("sqlserver://%s:%s@localhost:1433?database=%s", utilities.GoDotEnvVariable("UserID"), utilities.GoDotEnvVariable("Password"), utilities.GoDotEnvVariable("Database"))
	DbGorm, ErrGorm = gorm.Open("mssql", connectionString)

	os.Setenv("DBStatus", "dbcreateds")
	if utilities.GoDotEnvVariable("DBStatus") == "true" {
		//DbGorm.Debug().DropTableIfExists()
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_users{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_Registered{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_ImportationDocument{})

		//DbGorm.Debug().DropTableIfExists(&models.Tbl_KYC{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_role{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_status{})

		//DbGorm.Debug().DropTableIfExists(&models.Tbl_auditlog{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_clearing{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_goodstype{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_importation{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_servicetype{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_shipper{})
		//DbGorm.Debug().DropTableIfExists(&models.Tbl_subgoodstype{})

		DbGorm.SingularTable(true)
		//DbGorm.Debug().AutoMigrate(&models.Tbl_users{})
		//DbGorm.Debug().AutoMigrate(&models.Tbl_Registered{})
		//DbGorm.Debug().AutoMigrate(&models.Tbl_ImportationDocument{})
		//DbGorm.Debug().AutoMigrate(&models.Tbl_KYC{})
		//DbGorm.Debug().AutoMigrate(&models.Tbl_role{})
		//DbGorm.Debug().AutoMigrate(&models.Tbl_status{})

		// DbGorm.Debug().AutoMigrate(&models.Tbl_auditlog{})
		// DbGorm.Debug().AutoMigrate(&models.Tbl_clearing{})
		// DbGorm.Debug().AutoMigrate(&models.Tbl_goodstype{})
		// DbGorm.Debug().AutoMigrate(&models.Tbl_importation{})
		// DbGorm.Debug().AutoMigrate(&models.Tbl_servicetype{})
		// DbGorm.Debug().AutoMigrate(&models.Tbl_shipper{})
		// DbGorm.Debug().AutoMigrate(&models.Tbl_subgoodstype{})
	}

	fmt.Println("03")
	fmt.Println(os.Getenv("DBStatus"))
	fmt.Println("04")
	if ErrGorm != nil {
		fmt.Printf(ErrGorm.Error())
		return
	}
}

func GetDB() *gorm.DB {
	return DbGorm
}

// func CreateLog() {
// 	dt := time.Now()
// 	filename := "log-" + dt.Format("01-02-2006") + ".txt"
// 	_, Ferr := os.Stat(filename)
// 	if os.IsNotExist(Ferr) {
// 		ret, err := os.Create(filename)
// 		if err.Error() != "" {
// 			fmt.Println("Unable to create logfile for today" + filename)
// 		}
// 		fmt.Println(ret)
// 		defer ret.Close()
// 		log.SetOutput(ret)
// 	}
// }

// func LogError(err error) {
// 	errLog := &log.Logger{}
// 	errLog.Println(err)
// }
