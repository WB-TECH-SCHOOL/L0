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
        "/api/order/{order_id}": {
            "get": {
                "description": "Get order by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get Order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order data",
                        "schema": {
                            "$ref": "#/definitions/dto.Order"
                        }
                    },
                    "404": {
                        "description": "Order with that ID not found",
                        "schema": {
                            "$ref": "#/definitions/responses.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Order": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "responses.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
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
