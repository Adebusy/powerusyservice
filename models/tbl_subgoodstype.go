package models

import "time"

type Tbl_subgoodstype struct {
	Id          int       `json:"id" validate:"omitempty"`
	Goodstypeid int       `json:"Goodstypeid" validate:"omitempty"`
	Name        string    `json:"Name" validate:"omitempty"`
	Dateadded   time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid    int       `json:"StatusId" validate:"omitempty"`
}

type SubgoodstypeIn struct {
	Goodstypeid int       `json:"Goodstypeid" validate:"omitempty"`
	Name        string    `json:"Name" validate:"omitempty"`
	Dateadded   time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid    int       `json:"StatusId" validate:"omitempty"`
}

type SubgoodstypeOut struct {
	Id          int       `json:"id" validate:"omitempty"`
	Goodstypeid int       `json:"Goodstypeid" validate:"omitempty"`
	Name        string    `json:"Name" validate:"omitempty"`
	Dateadded   time.Time `json:"Dateadded" validate:"omitempty"`
	Statusid    int       `json:"StatusId" validate:"omitempty"`
}
