package models

type tbl_User struct {
	Id          int    `json:"Id" validate:"omitempty"`
	Username    string `json:"Username" validate:"omitempty"`
	FirstName   string `json:"FirstName" validate:"omitempty"`
	LastName    string `json:"LastName" validate:"omitempty"`
	MiddleName  string `json:"MiddleName" validate:"omitempty"`
	PhoneNumber string `json:"PhoneNumber" validate:"omitempty"`
	DateAdded   string `json:"DateAdded" validate:"omitempty"`
	RoleID      int    `json:"RoleID" validate:"omitempty"`
	Status      bool   `json:"Status" validate:"omitempty"`
	Email       string `json:"Email" validate:"omitempty"`
	Password    string `json:"Password" validate:"omitempty"`
	AuthCode    string `json:"AuthCode" validate:"omitempty"`
}

type UserIn struct {
	Username    string `json:"Username" validate:"omitempty"`
	FirstName   string `json:"FirstName" validate:"omitempty"`
	LastName    string `json:"LastName" validate:"omitempty"`
	MiddleName  string `json:"MiddleName" validate:"omitempty"`
	PhoneNumber string `json:"PhoneNumber" validate:"omitempty"`
	Email       string `json:"Email" validate:"omitempty"`
	Password    string `json:"Password" validate:"omitempty"`
}

type UserOut struct {
	Username    string `json:"Username" validate:"omitempty"`
	FirstName   string `json:"FirstName" validate:"omitempty"`
	LastName    string `json:"LastName" validate:"omitempty"`
	MiddleName  string `json:"MiddleName" validate:"omitempty"`
	PhoneNumber string `json:"PhoneNumber" validate:"omitempty"`
	DateAdded   string `json:"DateAdded" validate:"omitempty"`
	RoleID      int    `json:"RoleID" validate:"omitempty"`
	Status      bool   `json:"Status" validate:"omitempty"`
	Email       string `json:"Email" validate:"omitempty"`
	Password    string `json:"Password" validate:"omitempty"`
	AuthCode    string `json:"AuthCode" validate:"omitempty"`
}
