package sqlserver

import (
	"fmt"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/jinzhu/gorm"
)

func NewUser(db *gorm.DB) IUserService {
	return &dbconnect{db}
}

type IUserService interface {
	GetUserByEmailandUserId(email string, userId int) (models.Tbl_users, error)
	GetUserByEmail(email string) (models.Tbl_users, error)
	GetUserDetailsByUsername(username string) models.Tbl_users
	CreateUser(newusr *models.Tbl_users) error
	GetUserDetailsByUsernameAndPassword(username, password string) models.UserOut
	GetAllUsers() []models.UserOut
	CheckEmailWithAuthCode(email, authcode string) models.Tbl_users
	UpdateUserStatusUponAuthCode(email, authcode string, roleId int) bool
	GetUserById(userId int) (models.Tbl_users, error)
}

func (db dbconnect) GetUserByEmail(email string) (models.Tbl_users, error) {
	user := models.Tbl_users{}
	retQuery := db.DbGorm.Table(`Tbl_users`).Where(`Email =?`, email).First(&user).Error
	if retQuery != nil {
		fmt.Println(retQuery.Error())
	}
	return user, retQuery
}

func (db dbconnect) GetUserById(userId int) (models.Tbl_users, error) {
	user := models.Tbl_users{}
	retQuery := db.DbGorm.Table(`Tbl_users`).Where(`id =?`, userId).First(&user).Error
	if retQuery != nil {
		fmt.Println(retQuery.Error())
	}
	return user, retQuery
}

func (db dbconnect) GetUserByEmailandUserId(email string, userId int) (models.Tbl_users, error) {
	user := models.Tbl_users{}
	retQuery := db.DbGorm.Table(`Tbl_users`).Where(`Email =? and Id=?`, email, userId).First(&user).Error
	if retQuery != nil {
		fmt.Println(retQuery)
	}
	return user, retQuery
}

func (ts dbconnect) GetUserDetailsByUsername(username string) models.Tbl_users {
	reponseUser := models.Tbl_users{}
	retQuery := ts.DbGorm.Table(`Tbl_users`).Where(`username =?`, username).First(&reponseUser).Error
	if retQuery != nil {
		if retQuery.Error() == "record not found" {
			return reponseUser
		}
		fmt.Println(retQuery)
		//log.Fatal(retQuery.Error())
	}
	return reponseUser
}

func (ts dbconnect) CreateUser(newusr *models.Tbl_users) error {
	retQuery := ts.DbGorm.Table("Tbl_users").Create(&newusr).Error
	if retQuery != nil {
		fmt.Println(retQuery)
		//log.Fatal(retQuery.Error())
	}
	return retQuery
}

func (ts dbconnect) GetUserDetailsByUsernameAndPassword(username, password string) models.UserOut {
	reponseUser := models.UserOut{}
	retQuery := ts.DbGorm.Table(`Tbl_users`).Where(`username=? and password=?`, username, password).First(&reponseUser).Error
	if retQuery != nil {
		fmt.Println(retQuery)
		//log.Fatal(retQuery.Error())
	}
	return reponseUser
}

func (ts dbconnect) GetAllUsers() []models.UserOut {
	users := []models.UserOut{}
	retQuery := ts.DbGorm.Table(`Tbl_users`).Find(&users).Error
	if retQuery != nil {
		fmt.Println(retQuery)
		//log.Fatal(retQuery.Error())
	}
	return users
}

func (ts dbconnect) CheckEmailWithAuthCode(email, authcode string) models.Tbl_users {
	user := models.Tbl_users{}
	retQuery := ts.DbGorm.Table("Tbl_users").Where("authcode=? and email=?", authcode, email).Find(&user).Error
	if retQuery != nil {
		fmt.Println(retQuery)
		//log.Fatal(retQuery.Error())
	}
	return user
}

func (ts dbconnect) UpdateUserStatusUponAuthCode(email, authcode string, roleId int) bool {
	retQuery := ts.DbGorm.Table("Tbl_users").Where("authcode=? and email=?", authcode, email).Update("roleid", roleId).Error
	if retQuery != nil {
		fmt.Println(retQuery)
		//log.Fatal(retQuery.Error())
		return false
	}
	return true
}
func (ts dbconnect) UploadCompanyDocument(request models.Tbl_ImportationDocument) (int, error) {
	if query := ts.DbGorm.Table("Tbl_ImportationDocument").Create(&request).Error; query != nil {
		fmt.Println(query)
		//log.Fatal(query.Error())
		return 0, query
	}
	return 1, nil
}
