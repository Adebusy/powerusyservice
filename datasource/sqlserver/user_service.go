package sqlserver

import (
	"fmt"
	"log"

	"github.com/Adebusy/powerusyservice/models"
	"github.com/jinzhu/gorm"
)

func NewUser(db *gorm.DB) IUserService {
	return &dbconnect{db}
}

type IUserService interface {
	GetUserByEmailandUserId(email string, userId int) (models.Tbl_users, error)
	GetUserByEmail(email string) models.Tbl_users
	GetUserDetailsByUsername(username string) models.Tbl_users
	CreateNewUser(newusr *models.Tbl_users) error
	GetUserDetailsByUsernameAndPassword(username, password string) models.UserOut
	GetAllUsers() []models.UserOut
	CheckEmailWithAuthCode(email, authcode string) models.Tbl_users
}

func (db dbconnect) GetUserByEmail(email string) models.Tbl_users {
	user := models.Tbl_users{}
	if retQuery := db.DbGorm.Table(`Tbl_users`).Where(`Email =?`, email).First(&user).Error; retQuery != nil {
		fmt.Println(retQuery)
	}
	return user
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
	if retQuery.Error() != "" {
		log.Fatal(retQuery.Error())
	}
	return reponseUser
}

func (ts dbconnect) CreateNewUser(newusr *models.Tbl_users) error {
	retQuery := ts.DbGorm.Table("Tbl_users").Create(&newusr).Error
	if retQuery.Error() != "" {
		log.Fatal(retQuery.Error())
	}
	return retQuery
}

func (ts dbconnect) GetUserDetailsByUsernameAndPassword(username, password string) models.UserOut {
	reponseUser := models.UserOut{}
	retQuery := ts.DbGorm.Table(`Tbl_users`).Where(`username=? and password=?`, username, password).First(&reponseUser).Error
	//retQuery := ts.DbGorm.Table(`Tbl_users`).Where(`username =?`, username).First(&reponseUser).Error
	if retQuery.Error() != "" {
		log.Fatal(retQuery.Error())
	}
	return reponseUser
}

func (ts dbconnect) GetAllUsers() []models.UserOut {
	users := []models.UserOut{}
	retQuery := ts.DbGorm.Debug().Table(`Tbl_users`).Find(&users).Error
	if retQuery != nil {
		fmt.Println(retQuery)
		//util.LogError(retQuery)
		log.Fatal(retQuery.Error())
	}
	return users
}

func (ts dbconnect) CheckEmailWithAuthCode(email, authcode string) models.Tbl_users {
	user := models.Tbl_users{}
	retQuery := ts.DbGorm.Table("Tbl_users").Where("authcode=? and email=?", authcode, email).Find(&user).Error
	if retQuery != nil {
		fmt.Println(retQuery)
		log.Fatal(retQuery.Error())
	}
	return user
}
