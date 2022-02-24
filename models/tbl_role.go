package models

import "time"

type tbl_role struct {
	Id          int
	Code        string
	Name        string
	DateCreated time.Time
	Status      bool
}

type roleIn struct {
	Name   string
	Status bool
}

type roleOut struct {
	Id          int
	Code        string
	Name        string
	DateCreated time.Time
	Status      bool
}
