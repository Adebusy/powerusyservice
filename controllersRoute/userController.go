package controllersroute

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/swaggo/swag/example/celler/httputil"
)

type userService struct {
	DbGorm *gorm.DB
}

func (ts userService) GetUserDetail(ctx *gin.Context) (int, error) {
	userObj := &models.UserIn{}

	if err := ctx.ShouldBindJSON(userObj); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return 1, err
	}

	if userObj.Email == "" {
		ctx.JSON(http.StatusBadRequest, "Email address is required for this request")
		return
	}

	query := fmt.Sprintf("insert into tbl_User(Username, FirstName, LastName, MiddleName,PhoneNumber, DateAdded,RoleID,Status,Email,Password,AuthCode) values(%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)", userObj.Username, userObj.FirstName, userObj.LastName, userObj.MiddleName, userObj.PhoneNumber, time.Now(), "1", "true", userObj.Email, userObj.Password, "323223")
	inst := ts.DbGorm.Exec(query)
}

func CreateUser(ctx *gin.Context) {

}

func LoginUSer(ctx *gin.Context) {

}
