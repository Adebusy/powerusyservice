package sqlserver

import (
	"testing"

	"github.com/Adebusy/powerusyservice/app/driver"
)

var newuserService = NewUser(driver.GetDB())

func TestUpdateUserStatusUponAuthCode(t *testing.T) {
	if !newuserService.UpdateUserStatusUponAuthCode("test@gmail.com", "1234", 1) {
		t.Error("Failed UpdateUserStatusUponAuthCode")
	}
}

///Users/alaonr/go/src/github.com/Adebusy/powerusyservice/powerusyservice
