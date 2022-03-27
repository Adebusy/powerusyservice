// Package powerusyservice GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package powerusyservice

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Alao Ramon, Nabil Abubakar",
            "url": "http://www.swagger.io/support",
            "email": "alao.adebusy@gmail.com, nabbo247@yahoo.com"
        },
        "license": {
            "name": "AddyTech Solution Ltd",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Gets the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Shows the new status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/company/ApproveKYC": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Approve KYC document.",
                "parameters": [
                    {
                        "description": "KYC approval",
                        "name": "KYC",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.KYCApprovalIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/api/company/GetAllKYC": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Gets all companies KYC details.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.KYCsOut"
                            }
                        }
                    }
                }
            }
        },
        "/api/company/GetCompanyByCompanyName/{CompanyName}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Gets company details by company name.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CompanyDetailsOut"
                        }
                    }
                }
            }
        },
        "/api/company/GetCompanyDetailByEmailandCompName/{email}/{companyname}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Gets company details by email address and company name.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CompanyDetailsOut"
                        }
                    }
                }
            }
        },
        "/api/company/GetKYCbyCompanyId/{Id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Gets company KYC details by company Id.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.KYCOut"
                        }
                    }
                }
            }
        },
        "/api/company/RegisterCompany": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Register new company",
                "parameters": [
                    {
                        "description": "Company details",
                        "name": "company",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CompanyDetailsIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/api/company/UploadCompanyDocuments": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Upload company document",
                "parameters": [
                    {
                        "description": "Upload company document",
                        "name": "CompanyDoc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ImportationDocumentIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/api/company/UploadKYC": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Upload KYC document.",
                "parameters": [
                    {
                        "description": "KYC document",
                        "name": "KYC",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.KYCIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/api/shippers/ApproveShippersDocument": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shipper"
                ],
                "summary": "Approve shippers document",
                "parameters": [
                    {
                        "description": "User Details",
                        "name": "shipperDoc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ShipperApprovalReqIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/api/shippers/GetAllShippersDocument": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shipper"
                ],
                "summary": "Gets all shippers document",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ShipperOut"
                            }
                        }
                    }
                }
            }
        },
        "/api/shippers/GetShippersDocumentByCompanyname/{comapanyname}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shipper"
                ],
                "summary": "Gets shipper's document by company name",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.ShipperOut"
                            }
                        }
                    }
                }
            }
        },
        "/api/shippers/UploadShippersDocument": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "shipper"
                ],
                "summary": "Upload shippers document",
                "parameters": [
                    {
                        "description": "User Details",
                        "name": "shipperDoc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ShipperIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/api/users/CheckEmailWithAuthCode/{email}/{authcode}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "validates email with authentication code",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseMessage"
                        }
                    }
                }
            }
        },
        "/api/users/CreateNewUser": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Creates new user",
                "parameters": [
                    {
                        "description": "User Details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserOut"
                        }
                    }
                }
            }
        },
        "/api/users/GetAllUsers": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "gets list of all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.UserOut"
                            }
                        }
                    }
                }
            }
        },
        "/api/users/GetUserDetailsByEmail/{email}": {
            "get": {
                "description": "gets user details by email address.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "gets user details by email address.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tbl_users"
                        }
                    }
                }
            }
        },
        "/api/users/GetUserDetailsByUsername/{username}": {
            "get": {
                "description": "Gets user details by username.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Gets user details by username.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Tbl_users"
                        }
                    }
                }
            }
        },
        "/api/users/Login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "log user into the syste",
                "parameters": [
                    {
                        "description": "Login user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserOut"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CompanyDetailsIn": {
            "type": "object",
            "properties": {
                "AccountNumber": {
                    "type": "string"
                },
                "Approvalcomment": {
                    "type": "string"
                },
                "BankName": {
                    "type": "string"
                },
                "CompanyAddress": {
                    "type": "string"
                },
                "CompanyLocation": {
                    "type": "string"
                },
                "CompanyLogo": {
                    "type": "string"
                },
                "CompanyName": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "Email": {
                    "type": "string"
                },
                "PostAddress": {
                    "type": "string"
                },
                "ServiceId": {
                    "type": "integer"
                },
                "StatusId": {
                    "type": "integer"
                },
                "UserId": {
                    "type": "integer"
                },
                "WorkingDays": {
                    "type": "string"
                },
                "WorkingHours": {
                    "type": "string"
                }
            }
        },
        "models.CompanyDetailsOut": {
            "type": "object",
            "properties": {
                "AccountNumber": {
                    "type": "string"
                },
                "Approvalcomment": {
                    "type": "string"
                },
                "ApprovedBy": {
                    "type": "integer"
                },
                "BankName": {
                    "type": "string"
                },
                "CompanyAddress": {
                    "type": "string"
                },
                "CompanyLocation": {
                    "type": "string"
                },
                "CompanyLogo": {
                    "type": "string"
                },
                "CompanyName": {
                    "type": "string"
                },
                "Dateadded": {
                    "type": "string"
                },
                "Dateapproved": {
                    "type": "string"
                },
                "Description": {
                    "type": "string"
                },
                "PostAddress": {
                    "type": "string"
                },
                "ServiceId": {
                    "type": "integer"
                },
                "StatusId": {
                    "type": "integer"
                },
                "UserId": {
                    "type": "integer"
                },
                "WorkingDays": {
                    "type": "string"
                },
                "WorkingHours": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.ImportationDocumentIn": {
            "type": "object",
            "properties": {
                "DocumentName": {
                    "type": "string"
                },
                "DocumentPath": {
                    "type": "string"
                },
                "ImportationId": {
                    "type": "string"
                }
            }
        },
        "models.KYCApprovalIn": {
            "type": "object",
            "properties": {
                "Approvalcomment": {
                    "type": "string"
                },
                "Approvedby": {
                    "type": "integer"
                },
                "Registeredid": {
                    "type": "integer"
                },
                "Status": {
                    "type": "integer"
                }
            }
        },
        "models.KYCIn": {
            "type": "object",
            "properties": {
                "Articlesofassociation": {
                    "type": "string"
                },
                "Auditedfinancialstatement": {
                    "type": "string"
                },
                "Certificateofincorporation": {
                    "type": "string"
                },
                "Memorandomofassociation": {
                    "type": "string"
                },
                "Powerofattorneygranted": {
                    "type": "string"
                },
                "Registeredid": {
                    "type": "integer"
                },
                "Taxclearancecertificate": {
                    "type": "string"
                },
                "Validbusinesslicense": {
                    "type": "string"
                }
            }
        },
        "models.KYCOut": {
            "type": "object",
            "properties": {
                "Approvalcomment": {
                    "type": "string"
                },
                "Approvedby": {
                    "type": "integer"
                },
                "Articlesofassociation": {
                    "type": "string"
                },
                "Auditedfinancialstatement": {
                    "type": "string"
                },
                "Certificateofincorporation": {
                    "type": "string"
                },
                "Dateadded": {
                    "type": "string"
                },
                "Dateapproved": {
                    "type": "string"
                },
                "Memorandomofassociation": {
                    "type": "string"
                },
                "Powerofattorneygranted": {
                    "type": "string"
                },
                "Registeredid": {
                    "type": "integer"
                },
                "Statusid": {
                    "type": "integer"
                },
                "Taxclearancecertificate": {
                    "type": "string"
                },
                "Validbusinesslicense": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.KYCsOut": {
            "type": "object",
            "properties": {
                "Approvalcomment": {
                    "type": "string"
                },
                "Approvedby": {
                    "type": "integer"
                },
                "Articlesofassociation": {
                    "type": "string"
                },
                "Auditedfinancialstatement": {
                    "type": "string"
                },
                "Certificateofincorporation": {
                    "type": "string"
                },
                "Companyname": {
                    "type": "string"
                },
                "Dateadded": {
                    "type": "string"
                },
                "Dateapproved": {
                    "type": "string"
                },
                "Memorandomofassociation": {
                    "type": "string"
                },
                "Powerofattorneygranted": {
                    "type": "string"
                },
                "Registeredid": {
                    "type": "integer"
                },
                "Statusid": {
                    "type": "integer"
                },
                "Taxclearancecertificate": {
                    "type": "string"
                },
                "Validbusinesslicense": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.LoginIn": {
            "type": "object",
            "properties": {
                "Password": {
                    "type": "string"
                },
                "Username": {
                    "type": "string"
                }
            }
        },
        "models.ResponseMessage": {
            "type": "object",
            "properties": {
                "responseCode": {
                    "type": "string"
                },
                "responseMessage": {
                    "type": "string"
                }
            }
        },
        "models.ShipperApprovalReqIn": {
            "type": "object",
            "properties": {
                "ApprovedBy": {
                    "type": "integer"
                },
                "Comment": {
                    "type": "string"
                },
                "CompanyName": {
                    "type": "string"
                },
                "Statusid": {
                    "type": "integer"
                },
                "UserId": {
                    "type": "integer"
                }
            }
        },
        "models.ShipperIn": {
            "type": "object",
            "properties": {
                "CompanyName": {
                    "type": "string"
                },
                "Location": {
                    "type": "string"
                },
                "Phonenumber": {
                    "type": "string"
                },
                "UserId": {
                    "type": "integer"
                },
                "tin": {
                    "type": "string"
                },
                "tinpassword": {
                    "type": "string"
                }
            }
        },
        "models.ShipperOut": {
            "type": "object",
            "properties": {
                "ApprovedBy": {
                    "type": "integer"
                },
                "Comment": {
                    "type": "string"
                },
                "CompanyName": {
                    "type": "string"
                },
                "Location": {
                    "type": "string"
                },
                "Phonenumber": {
                    "type": "string"
                },
                "Statusid": {
                    "type": "string"
                },
                "UserId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "tin": {
                    "type": "string"
                },
                "tinpassword": {
                    "type": "string"
                }
            }
        },
        "models.Tbl_users": {
            "type": "object",
            "properties": {
                "authcode": {
                    "type": "string"
                },
                "dateAdded": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "middlename": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phonenumber": {
                    "type": "string"
                },
                "roleid": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UserIn": {
            "type": "object",
            "properties": {
                "Email": {
                    "type": "string"
                },
                "Firstname": {
                    "type": "string"
                },
                "Lastname": {
                    "type": "string"
                },
                "Middlename": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                },
                "Phonenumber": {
                    "type": "string"
                },
                "Roleid": {
                    "type": "integer"
                },
                "Username": {
                    "type": "string"
                }
            }
        },
        "models.UserOut": {
            "type": "object",
            "properties": {
                "Accountnumber": {
                    "type": "string"
                },
                "Approvedby": {
                    "type": "string"
                },
                "Dateapproved": {
                    "type": "string"
                },
                "Statusid": {
                    "type": "string"
                },
                "authcode": {
                    "type": "string"
                },
                "dateadded": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "middlename": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phonenumber": {
                    "type": "string"
                },
                "roleid": {
                    "type": "integer"
                },
                "status": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8060",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Powerusy backend service",
	Description:      "This is a backend web service created for Powerusy operations.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
