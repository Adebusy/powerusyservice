package models

type Tbl_shipper struct {
	Id          int    `json:"id" validate:"omitempty"`
	Userid      int    `json:"UserId" validate:"omitempty"`
	Companyname string `json:"CompanyName" validate:"omitempty"`
	Tin         string `json:"tin" validate:"omitempty"`
	Tinpassword string `json:"tinpassword" validate:"omitempty"`
	Location    string `json:"Location" validate:"omitempty"`
	Phonenumber string `json:"Phonenumber" validate:"omitempty"`
	Statusid    string `json:"Statusid" validate:"omitempty"`
	Approvedby  int    `json:"ApprovedBy" validate:"omitempty"`
	Comment     string `json:"Comment" validate:"omitempty"`
}

type ShipperIn struct {
	Userid      int    `json:"UserId" validate:"omitempty"`
	Companyname string `json:"CompanyName" validate:"omitempty"`
	Tin         string `json:"tin" validate:"omitempty"`
	Tinpassword string `json:"tinpassword" validate:"omitempty"`
	Location    string `json:"Location" validate:"omitempty"`
	Phonenumber string `json:"Phonenumber" validate:"omitempty"`
}

type ShipperOut struct {
	Id          int    `json:"id" validate:"omitempty"`
	Userid      int    `json:"UserId" validate:"omitempty"`
	Companyname string `json:"CompanyName" validate:"omitempty"`
	Tin         string `json:"tin" validate:"omitempty"`
	Tinpassword string `json:"tinpassword" validate:"omitempty"`
	Location    string `json:"Location" validate:"omitempty"`
	Phonenumber string `json:"Phonenumber" validate:"omitempty"`
	Statusid    string `json:"Statusid" validate:"omitempty"`
	Approvedby  int    `json:"ApprovedBy" validate:"omitempty"`
	Comment     string `json:"Comment" validate:"omitempty"`
}

type ShipperApprovalReqIn struct {
	Userid      int    `json:"UserId" validate:"omitempty"`
	Companyname string `json:"CompanyName" validate:"omitempty"`
	Statusid    int    `json:"Statusid" validate:"omitempty"`
	Approvedby  int    `json:"ApprovedBy" validate:"omitempty"`
	Comment     string `json:"Comment" validate:"omitempty"`
}
