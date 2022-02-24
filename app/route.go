package app

import (
	cont "github.com/Adebusy/powerusyservice/controllersRoute"
	"github.com/gin-gonic/gin"
)

func (app *App) RegisterRoutes() {
	svc := gin.Default()
	svc.GET("/", cont.GetUserDetail)
	svc.Run(":8055")
}
