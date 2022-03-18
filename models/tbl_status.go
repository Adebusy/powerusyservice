package models

import "time"

type Tbl_status struct {
	Id        int       `json:"Id" validate:"omitempty"`
	Name      string    `json:"Name" validate:"omitempty"`
	IsDeleted bool      `json:"IsDeleted" validate:"omitempty"`
	Dateadded time.Time `json:"Dateadded" validate:"omitempty"`
}

type StatusIn struct {
	Name      string `json:"Name" validate:"omitempty"`
	IsDeleted bool   `json:"IsDeleted" validate:"omitempty"`
}

type StatusOut struct {
	Id        int       `json:"Id" validate:"omitempty"`
	Name      string    `json:"Name" validate:"omitempty"`
	IsDeleted bool      `json:"IsDeleted" validate:"omitempty"`
	Dateadded time.Time `json:"Dateadded" validate:"omitempty"`
}
