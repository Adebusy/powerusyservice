package driver

import (
	"fmt"

	"github.com/Adebusy/powerusyservice/utilities"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DbGorm *gorm.DB
var ErrGorm error

func init() {
	godotenv.Load()
	DbGorm, ErrGorm = gorm.Open("mssql", "sqlserver://%s:%s@localhost:1433?database=%s", utilities.GoDotEnvVariable("UserID"), utilities.GoDotEnvVariable("Password"), utilities.GoDotEnvVariable("Database"))
	if ErrGorm != nil {
		fmt.Printf(ErrGorm.Error())
		return
	}
}

func GetDB() *gorm.DB {
	return DbGorm
}
