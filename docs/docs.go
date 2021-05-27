// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Alejandro Gabriel Guerrero",
            "url": "http://github.com/gbrayhan",
            "email": "gbrayhan@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/medicine/get-all": {
            "get": {
                "description": "Get all Medicines on the system",
                "tags": [
                    "medicine"
                ],
                "summary": "Get all Medicines",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Medicine"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.GeneralResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/medicine/get-by-id/{medicine_id}": {
            "get": {
                "description": "Get Medicines by ID on the system",
                "tags": [
                    "medicine"
                ],
                "summary": "Get medicines by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of medicine",
                        "name": "medicine_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Medicine"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.GeneralResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.GeneralResponse"
                        }
                    }
                }
            }
        },
        "/medicine/new": {
            "post": {
                "description": "Create new medicine on the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "medicine"
                ],
                "summary": "Create New Medicine",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.MedicineController"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Medicine"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.GeneralResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.GeneralResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.GeneralResponse": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "error description",
                        "other error description"
                    ]
                }
            }
        },
        "controllers.MedicineController": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Something"
                },
                "ean_code": {
                    "type": "string",
                    "example": "122000000021"
                },
                "laboratory": {
                    "type": "string",
                    "example": "Roche"
                },
                "name": {
                    "type": "string",
                    "example": "Paracetamol"
                }
            }
        },
        "models.Medicine": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2021-02-24 20:19:39"
                },
                "description": {
                    "type": "string",
                    "example": "Some Description"
                },
                "ean_code": {
                    "type": "string",
                    "example": "9900000124"
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "laboratory": {
                    "type": "string",
                    "example": "Roche"
                },
                "name": {
                    "type": "string",
                    "example": "Paracetamol"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2021-02-24 20:19:39"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "Boilerplate Golang",
	Description: "Documentation's Boilerplate Golang",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
