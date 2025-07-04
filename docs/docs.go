// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/user": {
            "get": {
                "description": "Getting Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Name",
                        "name": "names",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "cursor (las id uuid)",
                        "name": "cursor",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "lastTimestamp",
                        "name": "lastTimestamp",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "collectionFormat": "csv",
                        "description": "Filter created_at Like min-max (example: 2025-06-11T08:28:51.400404Z)",
                        "name": "orders[created_at]",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.RequestUserDTO"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Sent data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RequestUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.SuccessResponseDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/get-by-email": {
            "post": {
                "description": "Getting user by Email and Password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "Sent data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RequestUserByEmailDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.RequestUserDTO"
                            }
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Getting user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.RequestUserDTO"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update all sent fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Put user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RequestUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.SuccessResponseDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Deleting user by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.RequestUserDTO"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update only sent fields",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Patch user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User id (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RequestUserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.UserItemResultDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.RequestUserByEmailDto": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "Some user email"
                }
            }
        },
        "user.RequestUserDTO": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                },
                "email": {
                    "type": "string",
                    "example": "Some user email"
                },
                "password": {
                    "type": "string",
                    "example": "Some user password"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2022-01-01T00:00:00Z"
                }
            }
        },
        "user.SuccessResponseDto": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "user.UserItemResultDto": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
