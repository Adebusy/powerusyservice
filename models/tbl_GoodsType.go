package models

import "time"

type tbl_GoodsType struct {
	Id        int
	Name      string
	DateAdded time.Time
	Status    bool
}

type GoodsTypeIn struct {
	Name   string
	Status bool
}

type GoodsTypeOut struct {
	Id        int
	Name      string
	DateAdded time.Time
	Status    bool
}
