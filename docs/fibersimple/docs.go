// Package fibersimple GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package fibersimple

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/healthcheck": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
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
        "/login": {
            "post": {
                "description": "user success to login then generate active token",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "root"
                ],
                "summary": "check user login",
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
        "/v1/location": {
            "get": {
                "description": "it takes user token and fetch user location from db.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "root"
                ],
                "summary": "get user location.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "it takes user token and fetch user location from db and updates to it.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "root"
                ],
                "summary": "updates user new location if exits location it updates.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "it takes user token and fetch user location from db if exits otherwise creates new.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "root"
                ],
                "summary": "add user new location if exits location it updates.",
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
        "/v1/nearest/user": {
            "post": {
                "description": "it give nearest 10 user.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "root"
                ],
                "summary": "get Get Nearest User.",
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
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Fiber Swagger API",
	Description:      "This about user location details",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
