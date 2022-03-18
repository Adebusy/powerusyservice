package models

import "time"

type Tbl_KYC struct {
	Id                         int       `json:"id" validate:"omitempty"`
	Registeredid               int       `json:"Registeredid" validate:"omitempty"`
	Certificateofincorporation string    `json:"Certificateofincorporation" validate:"omitempty"`
	Memorandomofassociation    string    `json:"Memorandomofassociation" validate:"omitempty"`
	Articlesofassociation      string    `json:"Articlesofassociation" validate:"omitempty"`
	Powerofattorneygranted     string    `json:"Powerofattorneygranted" validate:"omitempty"`
	Validbusinesslicense       string    `json:"Validbusinesslicense" validate:"omitempty"`
	Auditedfinancialstatement  string    `json:"Auditedfinancialstatement" validate:"omitempty"`
	Taxclearancecertificate    string    `json:"Taxclearancecertificate" validate:"omitempty"`
	Dateadded                  time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid                   int       `json:"Statusid" validate:"omitempty"`
	Approvedby                 int       `json:"Approvedby" validate:"omitempty"`
	Approvalcomment            string    `json:"Approvalcomment" validate:"omitempty"`
	Dateapproved               time.Time `json:"Dateapproved" validate:"omitempty"`
}

type KYCIn struct {
	Registeredid               int    `json:"Registeredid" validate:"omitempty"`
	Certificateofincorporation string `json:"Certificateofincorporation" validate:"omitempty"`
	Memorandomofassociation    string `json:"Memorandomofassociation" validate:"omitempty"`
	Articlesofassociation      string `json:"Articlesofassociation" validate:"omitempty"`
	Powerofattorneygranted     string `json:"Powerofattorneygranted" validate:"omitempty"`
	Validbusinesslicense       string `json:"Validbusinesslicense" validate:"omitempty"`
	Auditedfinancialstatement  string `json:"Auditedfinancialstatement" validate:"omitempty"`
	Taxclearancecertificate    string `json:"Taxclearancecertificate" validate:"omitempty"`
}

type KYCOut struct {
	Id                         int       `json:"id" validate:"omitempty"`
	Registeredid               int       `json:"Registeredid" validate:"omitempty"`
	Certificateofincorporation string    `json:"Certificateofincorporation" validate:"omitempty"`
	Memorandomofassociation    string    `json:"Memorandomofassociation" validate:"omitempty"`
	Articlesofassociation      string    `json:"Articlesofassociation" validate:"omitempty"`
	Powerofattorneygranted     string    `json:"Powerofattorneygranted" validate:"omitempty"`
	Validbusinesslicense       string    `json:"Validbusinesslicense" validate:"omitempty"`
	Auditedfinancialstatement  string    `json:"Auditedfinancialstatement" validate:"omitempty"`
	Taxclearancecertificate    string    `json:"Taxclearancecertificate" validate:"omitempty"`
	Dateadded                  time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid                   int       `json:"Statusid" validate:"omitempty"`
	Approvedby                 int       `json:"Approvedby" validate:"omitempty"`
	Approvalcomment            string    `json:"Approvalcomment" validate:"omitempty"`
	Dateapproved               time.Time `json:"Dateapproved" validate:"omitempty"`
}
