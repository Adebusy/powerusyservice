package models

import "time"

type Tbl_role struct {
	Id          int       `json:"id" validate:"omitempty"`
	Code        string    `json:"Code" validate:"omitempty"`
	Name        string    `json:"Name" validate:"omitempty"`
	Datecreated time.Time `json:"Datecreated" validate:"omitempty"`
	Status      bool      `json:"Status" validate:"omitempty"`
}

type RoleIn struct {
	Code string `json:"Code" validate:"omitempty"`
	Name string `json:"Name" validate:"omitempty"`
}

type RoleOut struct {
	Id          int       `json:"id" validate:"omitempty"`
	Code        string    `json:"Code" validate:"omitempty"`
	Name        string    `json:"Name" validate:"omitempty"`
	Datecreated time.Time `json:"Datecreated" validate:"omitempty"`
	Status      bool      `json:"Status" validate:"omitempty"`
}
