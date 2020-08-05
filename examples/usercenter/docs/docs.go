// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-08-05 08:02:09.871038 +0800 CST m=+0.052378021

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
        "/call/WXLoginHandler": {
            "put": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "LogicModule"
                ],
                "summary": "WXLoginHandler",
                "parameters": [
                    {
                        "type": "string",
                        "description": "appID",
                        "name": "appID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/call/:module/:cmd/:opt": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "http web 模块"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "module",
                        "name": "module",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "cmd",
                        "name": "cmd",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "opt",
                        "name": "opt",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "data",
                        "name": "data",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    }
                }
            }
        },
        "/v1/wxlogin": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "http web 模块"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "WXLoginReq",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.WXLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "http.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "http.WXLoginReq": {
            "type": "object",
            "properties": {
                "appid": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
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
	Version:     "1.0.2",
	Host:        "localhost:11080",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Nethopper UserCenter Server",
	Description: "This is a sample server Petstore server.",
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
