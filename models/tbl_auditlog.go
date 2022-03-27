package models

import "time"

type Tbl_auditlog struct {
	Id                 int
	UserId             int
	Operationperformed string
	Ipaddress          string
	Pagevisited        string
	Dateaccessed       time.Time
}

type AuditlogIn struct {
	UserId             int
	Operationperformed string
	Ipaddress          string
	Pagevisited        string
	Dateaccessed       time.Time
}

type AuditlogOut struct {
	UserId             int
	Operationperformed string
	Ipaddress          string
	Pagevisited        string
	Dateaccessed       time.Time
}
