package driver

import (
	"fmt"

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
	if ErrGorm != nil {
		fmt.Printf(ErrGorm.Error())
		return
	}
}

func GetDB() *gorm.DB {
	return DbGorm
}
