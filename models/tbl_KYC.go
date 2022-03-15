package models

import "time"

type Tbl_KYC struct {
	Id                         int       `json:"id" validate:"omitempty"`
	RegisteredId               int       `json:"StatusId" validate:"omitempty"`
	CertificateOfIncorporation string    `json:"Certificateofincorporation" validate:"omitempty"`
	MemorandomOfAssociation    string    `json:"Memorandomofassociation" validate:"omitempty"`
	ArticlesOfAssociation      string    `json:"Articlesofassociation" validate:"omitempty"`
	PowerOfAttorneyGranted     string    `json:"Powerofattorneygranted" validate:"omitempty"`
	ValidBusinessLicense       string    `json:"Validbusinesslicense" validate:"omitempty"`
	AuditedFinancialStatement  string    `json:"Auditedfinancialstatement" validate:"omitempty"`
	TaxClearanceCertificate    string    `json:"Taxclearancecertificate" validate:"omitempty"`
	DateAdded                  time.Time `json:"Dateadded" validate:"omitempty"`
	StatusId                   int       `json:"Statusid" validate:"omitempty"`
	Approvedby                 int       `json:"Approvedby" validate:"omitempty"`
	ApprovalComment            string    `json:"Approvalcomment" validate:"omitempty"`
	DateApproved               time.Time `json:"Dateapproved" validate:"omitempty"`
}

type KYCIn struct {
	RegisteredId               int       `json:"StatusId" validate:"omitempty"`
	CertificateOfIncorporation string    `json:"Certificateofincorporation" validate:"omitempty"`
	MemorandomOfAssociation    string    `json:"Memorandomofassociation" validate:"omitempty"`
	ArticlesOfAssociation      string    `json:"Articlesofassociation" validate:"omitempty"`
	PowerOfAttorneyGranted     string    `json:"Powerofattorneygranted" validate:"omitempty"`
	ValidBusinessLicense       string    `json:"Validbusinesslicense" validate:"omitempty"`
	AuditedFinancialStatement  string    `json:"Auditedfinancialstatement" validate:"omitempty"`
	TaxClearanceCertificate    string    `json:"Taxclearancecertificate" validate:"omitempty"`
	DateAdded                  time.Time `json:"Dateadded" validate:"omitempty"`
	StatusId                   int       `json:"Statusid" validate:"omitempty"`
	Approvedby                 int       `json:"Approvedby" validate:"omitempty"`
	ApprovalComment            string    `json:"Approvalcomment" validate:"omitempty"`
	DateApproved               time.Time `json:"Dateapproved" validate:"omitempty"`
}

type KYCOut struct {
	RegisteredId               int       `json:"StatusId" validate:"omitempty"`
	CertificateOfIncorporation string    `json:"Certificateofincorporation" validate:"omitempty"`
	MemorandomOfAssociation    string    `json:"Memorandomofassociation" validate:"omitempty"`
	ArticlesOfAssociation      string    `json:"Articlesofassociation" validate:"omitempty"`
	PowerOfAttorneyGranted     string    `json:"Powerofattorneygranted" validate:"omitempty"`
	ValidBusinessLicense       string    `json:"Validbusinesslicense" validate:"omitempty"`
	AuditedFinancialStatement  string    `json:"Auditedfinancialstatement" validate:"omitempty"`
	TaxClearanceCertificate    string    `json:"Taxclearancecertificate" validate:"omitempty"`
	DateAdded                  time.Time `json:"Dateadded" validate:"omitempty"`
	StatusId                   int       `json:"Statusid" validate:"omitempty"`
	Approvedby                 int       `json:"Approvedby" validate:"omitempty"`
	ApprovalComment            string    `json:"Approvalcomment" validate:"omitempty"`
	DateApproved               time.Time `json:"Dateapproved" validate:"omitempty"`
}
