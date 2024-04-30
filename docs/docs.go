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
        "/issues": {
            "get": {
                "description": "Get a list of all issues",
                "produces": [
                    "application/json"
                ],
                "summary": "List all issues",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Issue"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new issue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new issue",
                "parameters": [
                    {
                        "description": "Issue object to be created",
                        "name": "issue",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Issue"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Issue"
                        }
                    }
                }
            }
        },
        "/issues/{id}": {
            "get": {
                "description": "Get an issue by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get an issue by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Issue ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Issue"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing issue",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update an existing issue",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Issue ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated issue object",
                        "name": "issue",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Issue"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Issue"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an issue by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete an issue by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Issue ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Issue": {
            "description": "Issue Structure",
            "type": "object",
            "properties": {
                "assigned_to": {
                    "description": "name of user who creates issue",
                    "type": "string"
                },
                "category": {
                    "description": "issue category",
                    "type": "string"
                },
                "created_by": {
                    "description": "name of user who creates issue",
                    "type": "string"
                },
                "created_on": {
                    "description": "name of user who creates issue",
                    "type": "string"
                },
                "eta": {
                    "description": "name of user who creates issue",
                    "type": "string"
                },
                "issue_description": {
                    "description": "issue descriptiom",
                    "type": "string"
                },
                "issue_id": {
                    "type": "integer"
                },
                "priority": {
                    "description": "name of user who creates issue",
                    "type": "string"
                },
                "state": {
                    "description": "issue state",
                    "type": "string"
                },
                "tags": {
                    "description": "name of user who creates issue",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updated_on": {
                    "description": "name of user who creates issue",
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
