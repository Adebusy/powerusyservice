package models

import "time"

type Tbl_clearing struct {
	Id                  int       `json:"Id" validate:"omitempty"`
	Userid              int       `json:"Userid" validate:"omitempty"`
	Goodsid             int       `json:"Goodsid" validate:"omitempty"`
	Finalinvoice        string    `json:"FinalInvoice" validate:"omitempty"`
	Billoflanding       string    `json:"Billoflanding" validate:"omitempty"`
	Packinglist         string    `json:"Packinglist" validate:"omitempty"`
	Combinedcertificate string    `json:"Combinedcertificate" validate:"omitempty"`
	Soncap              string    `json:"Soncap" validate:"omitempty"`
	Approvedbyid        string    `json:"Approvedbyid" validate:"omitempty"`
	Dateapproved        time.Time `json:"Dateapproved" validate:"omitempty"`
	Comment             string    `json:"Comment" validate:"omitempty"`
}

type ClearingIn struct {
	Userid              int       `json:"Userid" validate:"omitempty"`
	Goodsid             int       `json:"Goodsid" validate:"omitempty"`
	Finalinvoice        string    `json:"FinalInvoice" validate:"omitempty"`
	Billoflanding       string    `json:"Billoflanding" validate:"omitempty"`
	Packinglist         string    `json:"Packinglist" validate:"omitempty"`
	Combinedcertificate string    `json:"Combinedcertificate" validate:"omitempty"`
	Soncap              string    `json:"Soncap" validate:"omitempty"`
	Approvedbyid        string    `json:"Approvedbyid" validate:"omitempty"`
	Dateapproved        time.Time `json:"Dateapproved" validate:"omitempty"`
	Comment             string    `json:"Comment" validate:"omitempty"`
}

type ClearingOut struct {
	Id                  int       `json:"Id" validate:"omitempty"`
	Userid              int       `json:"Userid" validate:"omitempty"`
	Goodsid             int       `json:"Goodsid" validate:"omitempty"`
	Finalinvoice        string    `json:"FinalInvoice" validate:"omitempty"`
	Billoflanding       string    `json:"Billoflanding" validate:"omitempty"`
	Packinglist         string    `json:"Packinglist" validate:"omitempty"`
	Combinedcertificate string    `json:"Combinedcertificate" validate:"omitempty"`
	Soncap              string    `json:"Soncap" validate:"omitempty"`
	Approvedbyid        string    `json:"Approvedbyid" validate:"omitempty"`
	Dateapproved        time.Time `json:"Dateapproved" validate:"omitempty"`
	Comment             string    `json:"Comment" validate:"omitempty"`
}
