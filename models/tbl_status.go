package models

import "time"

type tbl_status struct {
	Id        int
	Name      string
	IsDeleted bool
	DateAdded time.Time
}

type statusIn struct {
	Name      string
	IsDeleted bool
}

type statusOut struct {
	Id        int
	Name      string
	IsDeleted bool
	DateAdded time.Time
}
