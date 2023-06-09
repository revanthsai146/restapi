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
    "paths":  "/employees": {
		"get": {
			"security": [
				{
					"ApiKeyAuth": []
				}
			],
			"consumes": [
				"application/json"
			],
			"tags": [
				"Employee"
			],
			"responses": {
				"200": {
					"description": "OK",
					"schema": {
						"type": "array",
						"items": {
							"$ref": "#/definitions/Employee"
						}
					}
				},
				"400": {
					"description": "Bad Request"
				},
				"404": {
					"description": "Not Found"
				},
				"405": {
					"description": "Method Not Allowed"
				},
				"500": {
					"description": "Internal Server Error"
				}
			}
		}
		"Employee": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "Name": {
                    "type": "string"
                },
                "Salary": {
                    "type": "integer"
                }
            }
        },
		"securityDefinitions": {
			"ApiKeyAuth": {
				"type": "apiKey",
				"name": "x-api-key",
				"in": "header"
			}
		}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Sample Get API",
	Description:      "This is my swagger api spec",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
	
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}