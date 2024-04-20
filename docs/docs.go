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
    "paths": {
        "/api/v1/auth/login": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Login login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Login user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.LoginBodyRes"
                        }
                    },
                    "400": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "GetProfileUser getProfileUser user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Get profile user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.UserProfileBodyRes"
                        }
                    },
                    "403": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Authorization": []
                    }
                ],
                "description": "UpdateProfileUser updateProfileUser user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Update profile user",
                "parameters": [
                    {
                        "description": "Update Profile User Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UpdateUserProfileBodyReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.UserProfileBodyRes"
                        }
                    },
                    "403": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Register userRegister user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Management"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "Register User Payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1.UserRegisterBodyReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.UserRegisterBodyRes"
                        }
                    },
                    "400": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "result",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1.LoginBodyRes": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/v1.LoginBodyUserRes"
                }
            }
        },
        "v1.LoginBodyUserRes": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "v1.UpdateUserProfileBodyReq": {
            "type": "object",
            "properties": {
                "fullname": {
                    "type": "string"
                },
                "phone_no": {
                    "type": "string"
                }
            }
        },
        "v1.UserProfileBodyRes": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/v1.UserProfileUserBodyRes"
                }
            }
        },
        "v1.UserProfileUserBodyRes": {
            "type": "object",
            "properties": {
                "fullname": {
                    "type": "string"
                },
                "phone_no": {
                    "type": "string"
                }
            }
        },
        "v1.UserRegisterBodyReq": {
            "type": "object",
            "properties": {
                "fullname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_no": {
                    "type": "string"
                }
            }
        },
        "v1.UserRegisterBodyRes": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/v1.UserRegisterUserBodyRes"
                }
            }
        },
        "v1.UserRegisterUserBodyRes": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Authorization": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "XForwardedForAuth": {
            "type": "apiKey",
            "name": "X-Forwarded-For",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "User Management - API",
	Description:      "Simple User Management Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
