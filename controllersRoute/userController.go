package controllersroute

import (
	"net/http"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

func GetUserDetail(ctx *gin.Context) {
	userObj := &models.UserIn{}

	if err := ctx.ShouldBindJSON(userObj); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

}
