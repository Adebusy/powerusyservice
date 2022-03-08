package driver

import (
	"fmt"
	"os"
	"time"

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

func CreateLog() {
	dt := time.Now()
	filename := "log-" + dt.Format("01-02-2006") + ".txt"
	_, Ferr := os.Stat(filename)
	if os.IsNotExist(Ferr) {
		_, err := os.Create(filename)
		if err.Error() != "" {
			fmt.Println("Unable to create logfile for today" + filename)
		}
	}
}
