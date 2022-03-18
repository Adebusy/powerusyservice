package models

import "time"

type Tbl_Registered struct {
	Id              int       `json:"id" validate:"omitempty"`
	Userid          int       `json:"UserId" validate:"omitempty"`
	Companyname     string    `json:"CompanyName" validate:"omitempty"`
	Serviceid       int       `json:"ServiceId" validate:"omitempty"`
	Description     string    `json:"Description" validate:"omitempty"`
	Companylocation string    `json:"CompanyLocation" validate:"omitempty"`
	Companyaddress  string    `json:"CompanyAddress" validate:"omitempty"`
	Companylogo     string    `json:"CompanyLogo" validate:"omitempty"`
	Postaddress     string    `json:"PostAddress" validate:"omitempty"`
	Workingdays     string    `json:"WorkingDays" validate:"omitempty"`
	Workinghours    string    `json:"WorkingHours" validate:"omitempty"`
	Bankname        string    `json:"BankName" validate:"omitempty"`
	Accountnumber   string    `json:"AccountNumber" validate:"omitempty"`
	Dateadded       time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid        int       `json:"StatusId" validate:"omitempty"`
	Approvedby      int       `json:"ApprovedBy" validate:"omitempty"`
	Dateapproved    time.Time `json:"Dateapproved" validate:"omitempty"`
	Approvalcomment string    `json:"Approvalcomment" validate:"omitempty"`
}

type CompanyDetailsOut struct {
	Id              int       `json:"id" validate:"omitempty"`
	Userid          int       `json:"UserId" validate:"omitempty"`
	Companyname     string    `json:"CompanyName" validate:"omitempty"`
	Serviceid       int       `json:"ServiceId" validate:"omitempty"`
	Description     string    `json:"Description" validate:"omitempty"`
	Companylocation string    `json:"CompanyLocation" validate:"omitempty"`
	Companyaddress  string    `json:"CompanyAddress" validate:"omitempty"`
	Companylogo     string    `json:"CompanyLogo" validate:"omitempty"`
	Postaddress     string    `json:"PostAddress" validate:"omitempty"`
	Workingdays     string    `json:"WorkingDays" validate:"omitempty"`
	Workinghours    string    `json:"WorkingHours" validate:"omitempty"`
	Bankname        string    `json:"BankName" validate:"omitempty"`
	Accountnumber   string    `json:"AccountNumber" validate:"omitempty"`
	Dateadded       time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid        int       `json:"StatusId" validate:"omitempty"`
	Approvedby      int       `json:"ApprovedBy" validate:"omitempty"`
	Dateapproved    time.Time `json:"Dateapproved" validate:"omitempty"`
	Approvalcomment string    `json:"Approvalcomment" validate:"omitempty"`
}

type CompanyDetailsIn struct {
	Email           string `json:"Email" validate:"omitempty"`
	UserId          int    `json:"UserId" validate:"omitempty"`
	CompanyName     string `json:"CompanyName" validate:"omitempty"`
	ServiceId       int    `json:"ServiceId" validate:"omitempty"`
	Description     string `json:"Description" validate:"omitempty"`
	CompanyLocation string `json:"CompanyLocation" validate:"omitempty"`
	CompanyAddress  string `json:"CompanyAddress" validate:"omitempty"`
	CompanyLogo     string `json:"CompanyLogo" validate:"omitempty"`
	PostAddress     string `json:"PostAddress" validate:"omitempty"`
	WorkingDays     string `json:"WorkingDays" validate:"omitempty"`
	WorkingHours    string `json:"WorkingHours" validate:"omitempty"`
	BankName        string `json:"BankName" validate:"omitempty"`
	AccountNumber   string `json:"AccountNumber" validate:"omitempty"`
	Statusid        int    `json:"StatusId" validate:"omitempty"`
	Approvalcomment string `json:"Approvalcomment" validate:"omitempty"`
}
type ServicesRendered struct {
	Documentation       string `json:"Documentation" validate:"omitempty"`
	ClearingAndTrucking string `json:"Clearingandtrucking" validate:"omitempty"`
}

type CompanyDocumentIns struct {
	Email           string
	UserId          int
	CompanyName     string
	ServiceId       int
	Description     string
	CompanyLocation string
	CompanyAddress  string
	CompanyLogo     string
	PostAddress     string
	WorkingDays     string
	WorkingHours    string
	BankName        string
	AccountNumber   string
}

type CompanyDocumentOuts struct {
	Userid          int       `json:"UserId" validate:"omitempty"`
	Companyname     string    `json:"CompanyName" validate:"omitempty"`
	Serviceid       int       `json:"ServiceId" validate:"omitempty"`
	Description     string    `json:"Description" validate:"omitempty"`
	Companylocation string    `json:"CompanyLocation" validate:"omitempty"`
	Companyaddress  string    `json:"CompanyAddress" validate:"omitempty"`
	Companylogo     string    `json:"CompanyLogo" validate:"omitempty"`
	Postaddress     string    `json:"PostAddress" validate:"omitempty"`
	Workingdays     string    `json:"WorkingDays" validate:"omitempty"`
	Workinghours    string    `json:"WorkingHours" validate:"omitempty"`
	Bankname        string    `json:"BankName" validate:"omitempty"`
	Accountnumber   string    `json:"AccountNumber" validate:"omitempty"`
	Dateadded       time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid        int       `json:"StatusId" validate:"omitempty"`
	Approvedby      int       `json:"ApprovedBy" validate:"omitempty"`
	Dateapproved    time.Time `json:"Dateapproved" validate:"omitempty"`
	Approvalcomment string    `json:"Approvalcomment" validate:"omitempty"`
}
