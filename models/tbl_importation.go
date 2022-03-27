package models

import "time"

type Tbl_importation struct {
	Id            int       `json:"Id" validate:"omitempty"`
	Shipperid     int       `json:"Shipperid" validate:"omitempty"`
	Goodstypeid   int       `json:"Goodstypeid" validate:"omitempty"`
	Subgoodtypeid string    `json:"Subgoodtypeid" validate:"omitempty"`
	Comment       string    `json:"Comment" validate:"omitempty"`
	Servicetypeid string    `json:"Servicetypeid" validate:"omitempty"`
	Status        int       `json:"Status" validate:"omitempty"`
	Dateadded     time.Time `json:"Dateadded" validate:"omitempty"`
}

type ImportationIn struct {
	Shipperid     int    `json:"Shipperid" validate:"omitempty"`
	Goodstypeid   int    `json:"Goodstypeid" validate:"omitempty"`
	Subgoodtypeid string `json:"Subgoodtypeid" validate:"omitempty"`
	Comment       string `json:"Comment" validate:"omitempty"`
	Servicetypeid string `json:"Servicetypeid" validate:"omitempty"`
	Status        int    `json:"Status" validate:"omitempty"`
}

type ImportationOut struct {
	Id            int       `json:"Id" validate:"omitempty"`
	Shipperid     int       `json:"Shipperid" validate:"omitempty"`
	Goodstypeid   int       `json:"Goodstypeid" validate:"omitempty"`
	Subgoodtypeid string    `json:"Subgoodtypeid" validate:"omitempty"`
	Comment       string    `json:"Comment" validate:"omitempty"`
	Servicetypeid string    `json:"Servicetypeid" validate:"omitempty"`
	Status        int       `json:"Status" validate:"omitempty"`
	Dateadded     time.Time `json:"Dateadded" validate:"omitempty"`
}
