package main

import (
	"database/sql"
	"net/http"

	"github.com/Adebusy/powerusyservice/database"
	"github.com/gin-gonic/gin"
)

var Db *sql.DB
var Err error

func init() {
	Database := &database.Database{}
	Db, Err = Database.LoadDVInstance()
}

func main() {

	svc := gin.Default()
	svc.GET("/", CheckService)

	svc.Run(":8059")

}

func CheckService(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "I am up and running") //fmt.Println("I am up and running now")
}
