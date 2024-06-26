// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://evil.com",
        "contact": {
            "name": "API Support",
            "url": "http:///evil.com",
            "email": "codewarrior666@mail.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/update": {
            "post": {
                "description": "Update information about ports in in-memory storage",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "description": "port data",
                        "name": "Input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PortData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "model.PortData": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "city": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "coordinates": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "country": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "regions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "timezone": {
                    "type": "string"
                },
                "unlocs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "JSONParser App API",
	Description:      "API server for parsing json files that store information about ports.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
