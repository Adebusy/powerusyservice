package models

type ResponseMessage struct {
	ResponseCode    string
	ResponseMessage string
}

type ConfigStruct struct {
	CreateTable          bool
	IsDropExistingTables bool
}
