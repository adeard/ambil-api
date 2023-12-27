// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "ADE ARDIAN",
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
        "/api/v1/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": " AuthRequest Schema ",
                        "name": "AuthRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/merchant": {
            "get": {
                "description": "Get All Merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "Get All Merchant",
                "parameters": [
                    {
                        "type": "string",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "coordinate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "is_active",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "phone_number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "picture",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "name": "score",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "updated_at",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.MerchantData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Create Merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "Create Merchant",
                "parameters": [
                    {
                        "description": " MerchantRequest Schema ",
                        "name": "MerchantRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.MerchantRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.MerchantData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/merchant/{merchant_id}": {
            "get": {
                "description": "Get Detail Merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "Get Detail Merchant",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " Merchant Id ",
                        "name": "merchant_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.MerchantData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Update Merchant",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "Update Merchant",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " Merchant Id ",
                        "name": "merchant_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " MerchantData Schema ",
                        "name": "MerchantData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.MerchantData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/order": {
            "get": {
                "description": "Get All Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get All Order",
                "parameters": [
                    {
                        "type": "string",
                        "name": "cancel_reason",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "created_at",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "merchant_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "notes",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order_no",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "order_status_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "payment_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "total_price",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "updated_at",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "user_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.OrderData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Create Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create Order",
                "parameters": [
                    {
                        "description": " OrderRequest Schema ",
                        "name": "OrderRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.OrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.OrderData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/order/{order_id}": {
            "get": {
                "description": "Get Detail Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Detail Order",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " Order Id ",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.OrderData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Update Order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update Order",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " Order Id ",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": " OrderData Schema ",
                        "name": "OrderData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.OrderData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/register": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": " RegisterRequest Schema ",
                        "name": "RegisterRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.UserData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.AuthRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.MerchantData": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "coordinate": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.MerchantRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "coordinate": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "is_active": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                },
                "score": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.OrderData": {
            "type": "object",
            "properties": {
                "cancel_reason": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "notes": {
                    "type": "string"
                },
                "order_date": {
                    "type": "string"
                },
                "order_no": {
                    "type": "string"
                },
                "order_status_id": {
                    "type": "integer"
                },
                "payment_date": {
                    "type": "string"
                },
                "total_price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.OrderRequest": {
            "type": "object",
            "properties": {
                "cancel_reason": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "merchant_id": {
                    "type": "integer"
                },
                "notes": {
                    "type": "string"
                },
                "order_date": {
                    "type": "string"
                },
                "order_no": {
                    "type": "string"
                },
                "order_status_id": {
                    "type": "integer"
                },
                "payment_date": {
                    "type": "string"
                },
                "total_price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.RegisterRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "fullname",
                "password",
                "phone_number"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "domain.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "elapsed_time": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.UserData": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:6969",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "API SWAGGER FOR AMBIL API SERVICE",
	Description:      "AMBIL API SERVICE",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
