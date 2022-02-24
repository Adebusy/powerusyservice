package main

import (
	"database/sql"
	"net/http"

	"github.com/Adebusy/powerusyservice/database"
	"github.com/gin-gonic/gin"
)

var db *sql.DB
var err error

func init() {
	Database := &database.Database{}
	db, err = Database.LoadDVInstance()
}

func main() {

	svc := gin.Default()
	svc.GET("/", CheckService)

	svc.Run(":8055")

}

func CheckService(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "I am up and running") //fmt.Println("I am up and running now")
}
