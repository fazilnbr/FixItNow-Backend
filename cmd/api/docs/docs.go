// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/add-profile": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Profile Management"
                ],
                "summary": "Add User Profile And Update Mail",
                "operationId": "AddProfileAndUpdateMail",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "userData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_domain.UserData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_utils.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_utils.Response"
                        }
                    }
                }
            }
        },
        "/user/login-gl": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "Authenticate With Google",
                "operationId": "Authenticate With Google",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_utils.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_utils.Response"
                        }
                    }
                }
            }
        },
        "/user/sent-otp": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "Send OTP for Users",
                "operationId": "sendOtp",
                "parameters": [
                    {
                        "description": "Mobile Number",
                        "name": "mobileNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_domain.Signup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_utils.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_utils.Response"
                        }
                    }
                }
            }
        },
        "/user/signup-and-login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Authentication"
                ],
                "summary": "SignUp for users",
                "operationId": "SignUp authentication",
                "parameters": [
                    {
                        "description": "Mobile Number And OTP",
                        "name": "mobileNumberAndOTP",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_domain.Signup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_utils.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_fazilnbr_project-workey_pkg_utils.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_fazilnbr_project-workey_pkg_domain.Signup": {
            "type": "object",
            "properties": {
                "countrycode": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                },
                "phonenumber": {
                    "type": "string"
                }
            }
        },
        "github_com_fazilnbr_project-workey_pkg_domain.UserData": {
            "type": "object",
            "required": [
                "profilephoto"
            ],
            "properties": {
                "dob": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "profilephoto": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "github_com_fazilnbr_project-workey_pkg_utils.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
