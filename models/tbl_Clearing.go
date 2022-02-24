package models

import "time"

type tbl_Clearing struct {
	Id                  int
	UserId              int
	GoodsId             int
	FinalInvoice        string
	BillOfLading        string
	PackingList         string
	CombinedCertificate string
	SONCAP              string
	ApprovedById        string
	DateApproved        time.Time
	Comment             string
}

type ClearingIn struct {
	UserId  int
	GoodsId int
}

type ClearingOut struct {
	UserId              int
	GoodsId             int
	FinalInvoice        string
	BillOfLading        string
	PackingList         string
	CombinedCertificate string
	SONCAP              string
	ApprovedById        string
	DateApproved        time.Time
	Comment             string
}
