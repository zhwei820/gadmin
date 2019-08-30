// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-30 17:55:48.1969796 +0800 CST m=+0.182994901

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
        "termsOfService": "http://hello.io/terms/",
        "contact": {
            "name": "hello",
            "url": "http://www.hello.io",
            "email": "hello@hello.io"
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
        "/rbac/login": {
            "post": {
                "description": "登陆",
                "tags": [
                    "auth"
                ],
                "summary": "登陆",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "Login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rbac/logout": {
            "get": {
                "description": "登出",
                "tags": [
                    "auth"
                ],
                "summary": "登出 （jwt 后端暂时无法登出）",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rbac/policy": {
            "get": {
                "description": "policy list",
                "tags": [
                    "policy"
                ],
                "summary": "policy list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "UpdatePolicy",
                "tags": [
                    "policy"
                ],
                "summary": "UpdatePolicy",
                "parameters": [
                    {
                        "description": "UpdatePolicy",
                        "name": "UpdatePolicy",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.UpdatePolicy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create policy",
                "tags": [
                    "policy"
                ],
                "summary": "create policy",
                "parameters": [
                    {
                        "description": "UpdatePolicy",
                        "name": "UpdatePolicy",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.UpdatePolicy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rbac/policy/byrole": {
            "get": {
                "description": "GetPolicyByRole",
                "tags": [
                    "policy"
                ],
                "summary": "GetPolicyByRole",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role",
                        "name": "role",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "SetPolicyByRole",
                "tags": [
                    "policy"
                ],
                "summary": "SetPolicyByRole",
                "parameters": [
                    {
                        "description": "SetPolicyByRole",
                        "name": "SetPolicyByRole",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.SetPolicyByRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rbac/refresh_token": {
            "get": {
                "description": "刷新token",
                "tags": [
                    "auth"
                ],
                "summary": "刷新token",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rbac/role": {
            "get": {
                "description": "role list",
                "tags": [
                    "role"
                ],
                "summary": "role list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update role",
                "tags": [
                    "role"
                ],
                "summary": "Update role",
                "parameters": [
                    {
                        "description": "PostRole",
                        "name": "PostRole",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.PostRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "create role",
                "tags": [
                    "role"
                ],
                "summary": "create role",
                "parameters": [
                    {
                        "description": "PostRole",
                        "name": "PostRole",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.PostRole"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete role",
                "tags": [
                    "role"
                ],
                "summary": "delete role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role",
                        "name": "role",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rbac/role/userrole": {
            "put": {
                "description": "SetRoleByUserName",
                "tags": [
                    "role"
                ],
                "summary": "SetRoleByUserName",
                "parameters": [
                    {
                        "description": "SetRoleByUserName",
                        "name": "SetRoleByUserName",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.SetRoleByUserName"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rbac/user": {
            "get": {
                "description": "user list",
                "tags": [
                    "user"
                ],
                "summary": "user list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "UpdateUser",
                "tags": [
                    "user"
                ],
                "summary": "UpdateUser",
                "parameters": [
                    {
                        "description": "UpdateUser",
                        "name": "UpdateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "CreateUser",
                "tags": [
                    "user"
                ],
                "summary": "CreateUser",
                "parameters": [
                    {
                        "description": "CreateUser",
                        "name": "CreateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api_model.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete user",
                "tags": [
                    "user"
                ],
                "summary": "delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/rbac/user/info": {
            "get": {
                "description": "user info",
                "tags": [
                    "user"
                ],
                "summary": "user info",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api_model.CreateUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "passwordconfirm": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api_model.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api_model.PostRole": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "api_model.SetPolicyByRole": {
            "type": "object",
            "properties": {
                "policys": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "api_model.SetRoleByUserName": {
            "type": "object",
            "properties": {
                "roles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "api_model.UpdatePolicy": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "api_model.UpdateUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "passwordconfirm": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Swagger Example API",
	Description: "This is a hello server .",
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
