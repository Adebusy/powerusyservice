package models

import "time"

type Tbl_goodstype struct {
	Id        int       `json:"Id" validate:"omitempty"`
	Name      string    `json:"Name" validate:"omitempty"`
	Dateadded time.Time `json:"Dateadded" validate:"omitempty"`
	Status    bool      `json:"Status" validate:"omitempty"`
}

type GoodsTypeIn struct {
	Name      string    `json:"Name" validate:"omitempty"`
	Dateadded time.Time `json:"Dateadded" validate:"omitempty"`
	Status    bool      `json:"Status" validate:"omitempty"`
}

type GoodsTypeOut struct {
	Id        int       `json:"Id" validate:"omitempty"`
	Name      string    `json:"Name" validate:"omitempty"`
	Dateadded time.Time `json:"Dateadded" validate:"omitempty"`
	Status    bool      `json:"Status" validate:"omitempty"`
}
