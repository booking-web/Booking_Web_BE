// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/api/v1/change-password": {
            "post": {
                "description": "Returns Change Password Successful",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Change Password",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "userLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ForgotPass"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "userId",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseDataType"
                        }
                    }
                }
            }
        },
        "/api/v1/get-user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Returns user details by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user details by Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Id",
                        "name": "userId",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Users"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "Returns Login Successful",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "userLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ResponseLogin"
                        }
                    }
                }
            }
        },
        "/api/v1/sign-up": {
            "post": {
                "description": "Returns  Sign Up Successful",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Sign Up Account",
                "parameters": [
                    {
                        "description": "New User",
                        "name": "userId",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.Users"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.ResponseDataType"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ForgotPass": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "handlers.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handlers.ResponseDataType": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        },
        "types.ResponseLogin": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "types.Users": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "roleId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
