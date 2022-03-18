package models

import "time"

type Tbl_ImportationDocument struct {
	Id              int       `json:"id" validate:"omitempty"`
	Importationid   string    `json:"importationId" validate:"omitempty"`
	Documentname    string    `json:"documentName" validate:"omitempty"`
	Documentpath    string    `json:"documentPath" validate:"omitempty"`
	Dateadded       time.Time `json:"dateadded" validate:"omitempty"`
	Statusid        int       `json:"statusId" validate:"omitempty"`
	Approvedby      string    `json:"approvedby" validate:"omitempty"`
	Dateapproved    string    `json:"dateapproved" validate:"omitempty"`
	Approvalcomment string    `json:"approvalcomment" validate:"omitempty"`
}
type ImportationDocumentIn struct {
	Importationid string `json:"ImportationId" validate:"omitempty"`
	Documentname  string `json:"DocumentName" validate:"omitempty"`
	Documentpath  string `json:"DocumentPath" validate:"omitempty"`
}

type ImportationDocumentOut struct {
	Id              int       `json:"Id" validate:"omitempty"`
	Importationid   string    `json:"ImportationId" validate:"omitempty"`
	Documentname    string    `json:"DocumentName" validate:"omitempty"`
	Documentpath    string    `json:"DocumentPath" validate:"omitempty"`
	Dateadded       time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid        int       `json:"StatusId" validate:"omitempty"`
	Approvedby      string    `json:"Approvedby" validate:"omitempty"`
	Dateapproved    string    `json:"DateApproved" validate:"omitempty"`
	Approvalcomment string    `json:"ApprovalComment" validate:"omitempty"`
}
