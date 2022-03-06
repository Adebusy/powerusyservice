package models

import "time"

type Tbl_users struct {
	Id          int       `json:"id" validate:"omitempty"`
	Firstname   string    `json:"firstname" validate:"omitempty"`
	Username    string    `json:"username" validate:"omitempty"`
	Lastname    string    `json:"lastname" validate:"omitempty"`
	Middlename  string    `json:"middlename" validate:"omitempty"`
	Phonenumber string    `json:"phonenumber" validate:"omitempty"`
	Dateadded   time.Time `json:"dateAdded" validate:"omitempty"`
	Roleid      int       `json:"roleid" validate:"omitempty"`
	Status      bool      `json:"status" validate:"omitempty"`
	Email       string    `json:"email" validate:"omitempty"`
	Password    string    `json:"password" validate:"omitempty"`
	Authcode    string    `json:"authcode" validate:"omitempty"`
}

// type Tbl_users struct {
// 	Id          int       `json:"Id" validate:"omitempty"`
// 	Username    string    `json:"username" validate:"omitempty"`
// 	Firstname   string    `json:"FirstName" validate:"omitempty"`
// 	Lastname    string    `json:"LastName" validate:"omitempty"`
// 	Middlename  string    `json:"MiddleName" validate:"omitempty"`
// 	Phonenumber string    `json:"PhoneNumber" validate:"omitempty"`
// 	Dateadded   time.Time `json:"DateAdded" validate:"omitempty"`
// 	Roleid      int       `json:"Roleid" validate:"omitempty"`
// 	Status      bool      `json:"Status" validate:"omitempty"`
// 	Email       string    `json:"Email" validate:"omitempty"`
// 	Password    string    `json:"Password" validate:"omitempty"`
// 	Authcode    string    `json:"AuthCode" validate:"omitempty"`
// }

type Tbl_user struct {
	Id        int    `json:"Id" validate:"omitempty"`
	FirstName string `json:"FirstName" validate:"omitempty"`
}

// type UserIn struct {
// 	Username    string `json:"Username" validate:"omitempty"`
// 	FirstName   string `json:"FirstName" validate:"omitempty"`
// 	LastName    string `json:"LastName" validate:"omitempty"`
// 	MiddleName  string `json:"MiddleName" validate:"omitempty"`
// 	PhoneNumber string `json:"PhoneNumber" validate:"omitempty"`
// 	Email       string `json:"Email" validate:"omitempty"`
// 	Password    string `json:"Password" validate:"omitempty"`
// }

type UserIn struct {
	Username    string `json:"Username" validate:"omitempty"`
	FirstName   string `json:"Firstname" validate:"omitempty"`
	LastName    string `json:"Lastname" validate:"omitempty"`
	MiddleName  string `json:"Middlename" validate:"omitempty"`
	PhoneNumber string `json:"Phonenumber" validate:"omitempty"`
	Email       string `json:"Email" validate:"omitempty"`
	Password    string `json:"Password" validate:"omitempty"`
}

type UserOut struct {
	Id          int       `json:"id" validate:"omitempty"`
	Firstname   string    `json:"firstname" validate:"omitempty"`
	Username    string    `json:"username" validate:"omitempty"`
	Lastname    string    `json:"lastname" validate:"omitempty"`
	Middlename  string    `json:"middlename" validate:"omitempty"`
	Phonenumber string    `json:"phonenumber" validate:"omitempty"`
	Dateadded   time.Time `json:"dateAdded" validate:"omitempty"`
	Roleid      int       `json:"roleid" validate:"omitempty"`
	Status      bool      `json:"status" validate:"omitempty"`
	Email       string    `json:"email" validate:"omitempty"`
	Password    string    `json:"password" validate:"omitempty"`
	Authcode    string    `json:"authcode" validate:"omitempty"`
}

type LoginIn struct {
	Username string `json:"Username" validate:"omitempty"`
	Password string `json:"Password" validate:"omitempty"`
}
