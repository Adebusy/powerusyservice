package models

import "time"

type tbl_serviceType struct {
	Id        int
	Service   string
	DateAdded time.Time
	isDeleted bool
}

type serviceTypeIn struct {
	Service string
}

type serviceTypeOut struct {
	Id        int
	Service   string
	DateAdded time.Time
	isDeleted bool
}
