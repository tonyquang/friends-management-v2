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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/add-friends": {
            "post": {
                "description": "Create a friend connection between two email addresses.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friendship"
                ],
                "summary": "Make Friend Connection",
                "parameters": [
                    {
                        "description": "RequestCreateUser",
                        "name": "friends",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/friendship.RequestFriend"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPError"
                        }
                    }
                }
            }
        },
        "/block": {
            "post": {
                "description": "Block updates from an email address.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friendship"
                ],
<<<<<<< HEAD
                "summary": "Block Subscribe to update an user",
=======
                "summary": "Block update an user",
>>>>>>> DONE-API
                "parameters": [
                    {
                        "description": "Requestor and Target to block update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/friendship.RequestUpdate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPError"
                        }
                    }
                }
            }
        },
        "/create-user": {
            "post": {
                "description": "Create A New User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create A New User",
                "parameters": [
                    {
                        "description": "RequestCreateUser",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RequestCreateUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPError"
                        }
                    }
                }
            }
        },
        "/get-list-friends": {
            "post": {
                "description": "Retrieve the friends list for an email address.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friendship"
                ],
                "summary": "Get Friends List",
                "parameters": [
                    {
                        "description": "RequestListFriends",
                        "name": "email",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/friendship.RequestListFriends"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/friendship.ResponeListFriends"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPError"
                        }
                    }
                }
            }
        },
        "/get-list-users-receive-update": {
            "post": {
                "description": "Retrieve all email addresses that can receive updates from an email address.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friendship"
                ],
                "summary": "Get Users Receive Update",
                "parameters": [
                    {
                        "description": "Sender and Text",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/friendship.RequestReceiveUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/friendship.ResponeReceiveUpdate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPError"
                        }
                    }
                }
            }
        },
        "/get-mutual-list-friends": {
            "post": {
                "description": "Retrieve the common friends list between two email addresses.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friendship"
                ],
                "summary": "Get Mutual Friends List",
                "parameters": [
                    {
                        "description": "RequestFriend",
                        "name": "friends",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/friendship.RequestFriend"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/friendship.ResponeListFriends"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPError"
                        }
                    }
                }
            }
        },
        "/list-users": {
            "get": {
                "description": "Get list users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "List users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.ResponeListUser"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPError"
                        }
                    }
                }
            }
        },
        "/subscribe": {
            "post": {
                "description": "Subscribe to updates from an email address.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Friendship"
                ],
                "summary": "Subscribe update an user",
                "parameters": [
                    {
                        "description": "Requestor and Target to subscribe update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/friendship.RequestUpdate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common_respone.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common_respone.HTTPError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "any error"
                }
            }
        },
        "common_respone.HTTPSuccess": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "friendship.RequestFriend": {
            "type": "object",
            "required": [
                "friends"
            ],
            "properties": {
                "friends": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "friendship.RequestListFriends": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "friendship.RequestReceiveUpdate": {
            "type": "object",
            "required": [
                "sender",
                "text"
            ],
            "properties": {
                "sender": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "friendship.RequestUpdate": {
            "type": "object",
            "required": [
                "requestor",
                "target"
            ],
            "properties": {
                "requestor": {
                    "type": "string"
                },
                "target": {
                    "type": "string"
                }
            }
        },
        "friendship.ResponeListFriends": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "friends": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "friendship.ResponeReceiveUpdate": {
            "type": "object",
            "properties": {
                "recipients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "user.RequestCreateUser": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "user.ResponeListUser": {
            "type": "object",
            "required": [
                "count",
                "list_users"
            ],
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list_users": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
