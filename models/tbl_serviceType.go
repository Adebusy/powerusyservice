package models

import "time"

type Tbl_servicetype struct {
	Id        int       `json:"Id" validate:"omitempty"`
	Service   string    `json:"Service" validate:"omitempty"`
	Isdeleted bool      `json:"Isdeleted" validate:"omitempty"`
	Dateadded time.Time `json:"Dateadded" validate:"omitempty"`
}

type ServicetypeIn struct {
	Service   string `json:"Service" validate:"omitempty"`
	Isdeleted bool   `json:"Isdeleted" validate:"omitempty"`
}

type ServicetypeOut struct {
	Id        int       `json:"Id" validate:"omitempty"`
	Service   string    `json:"Service" validate:"omitempty"`
	Isdeleted bool      `json:"Isdeleted" validate:"omitempty"`
	Dateadded time.Time `json:"Dateadded" validate:"omitempty"`
}
