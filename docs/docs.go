// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/activity": {
            "get": {
                "description": "Get a list of Activity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "Get All Activity.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : ' Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Create Activity .",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "Create Activity .",
                "parameters": [
                    {
                        "description": "the body to create a new Activity",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FormActivity"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : ' Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/activity/:id": {
            "get": {
                "description": "Get a Activity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "Get Activity By Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : ' Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "put": {
                "description": "Update Activity .",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "Update Activity .",
                "parameters": [
                    {
                        "description": "the body to Update a new Activity",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FormActivity"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : ' Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete a Activity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "Delete Activity By Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : ' Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/activity/riwayat": {
            "get": {
                "description": "Riwayat Activity.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Activity"
                ],
                "summary": "Riwayat Activity.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : ' Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/attendance": {
            "post": {
                "description": "Create Attendance .",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attendance"
                ],
                "summary": "Create Attendance .",
                "parameters": [
                    {
                        "description": "the body to create a new Attendance",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FormCreateAttendance"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : ' Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/attendance/riwayat": {
            "get": {
                "description": "Get a list of Riwayat Attendance.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attendance"
                ],
                "summary": "Riwayat Attendance.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization. How to input in swagger : ' Bearer \u003cinsert_your_token_here\u003e'",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Create User .",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AUTH"
                ],
                "summary": "Create User .",
                "parameters": [
                    {
                        "description": "the body to register user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FormRegister"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.FormRegister"
                        }
                    }
                }
            }
        },
        "/auth/token/": {
            "post": {
                "description": "Login User.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AUTH"
                ],
                "summary": "Login User.",
                "parameters": [
                    {
                        "description": "the body to login user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FormLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FormLogin"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.FormActivity": {
            "type": "object",
            "properties": {
                "activity": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                }
            }
        },
        "models.FormCreateAttendance": {
            "type": "object",
            "properties": {
                "type_attendance_id": {
                    "type": "integer"
                }
            }
        },
        "models.FormLogin": {
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
        "models.FormRegister": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
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