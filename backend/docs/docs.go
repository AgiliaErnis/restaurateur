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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/autocomplete": {
            "get": {
                "description": "Provides restaurant candidates for autocompletion based on provided input",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "autocomplete"
                ],
                "summary": "Autocomplete backend",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of searched restaurant",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "address of searched restaurant",
                        "name": "address",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseAutocompleteJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Logs in a user if the request headers contain an authenticated cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Logs in a user",
                "parameters": [
                    {
                        "description": "Logs in a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseUserJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logout"
                ],
                "summary": "Logs out a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/prague-college/restaurants": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "PC restaurants"
                ],
                "summary": "Returns restaurants around Prague College",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Radius (in meters) of the area around a provided or pre-selected starting point. Restaurants in this area will be returned. Radius can be ignored when specified with radius=ignore and lat and lon parameters will no longer be required. When no radius is provided, a default value of 1000 meters is used.",
                        "name": "radius",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filters restaurants based on a list of cuisines, separated by commas -\u003e cuisine=Czech,English. A restaurant will be returned only if it satisfies all provided cuisines.Available cuisines: American, Italian, Asian, Indian, Japanese, Vietnamese, Spanish, Mediterranean, French, Thai, Mexican, International, Czech, English, Balkan, Brazil, Russian, Chinese, Greek, Arabic, Korean.",
                        "name": "cuisine",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filters restaurants based on a list of price ranges, separated by commas -\u003e price-range=0-300,600-. A restaurant will be returned if it satisfies at least one provided price range. Available price ranges: 0-300,300-600,600-",
                        "name": "price-range",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all non vegetarian restaurants.",
                        "name": "vegetarian",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all non vegan restaurants.",
                        "name": "vegan",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all non gluten free restaurants.",
                        "name": "gluten-free",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all restaurants that don't have a takeaway option.",
                        "name": "takeaway",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all restaurants that don't have a delivery option.",
                        "name": "delivery-options",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sorts restaurants. Available sort options: price-asc, price-desc, rating",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseFullJSON"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "register"
                ],
                "summary": "Registers a user",
                "parameters": [
                    {
                        "description": "Create a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/db.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseUserJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/restaurant/{restaurant-id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurant"
                ],
                "summary": "Provides info about a specific restaurant",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Restaurant ID",
                        "name": "restaurant-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseFullJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/restaurants": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurants"
                ],
                "summary": "Returns restaurants based on queries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Radius (in meters) of the area around a provided or pre-selected starting point. Restaurants in this area will be returned. Radius can be ignored when specified with radius=ignore and lat and lon parameters will no longer be required. When no radius is provided, a default value of 1000 meters is used.",
                        "name": "radius",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Starting point for a search in a given radius.",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Latitude in degrees. Lat is required if radius is not set to ignore.",
                        "name": "lat",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "Longitude in degrees. Lon is required if radius is not set to ignore.",
                        "name": "lon",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filters restaurants based on a list of cuisines, separated by commas -\u003e cuisine=Czech,English. A restaurant will be returned only if it satisfies all provided cuisines.Available cuisines: American, Italian, Asian, Indian, Japanese, Vietnamese, Spanish, Mediterranean, French, Thai, Mexican, International, Czech, English, Balkan, Brazil, Russian, Chinese, Greek, Arabic, Korean.",
                        "name": "cuisine",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filters restaurants based on a list of price ranges, separated by commas -\u003e price-range=0-300,600-. A restaurant will be returned if it satisfies at least one provided price range. Available price ranges: 0-300,300-600,600-",
                        "name": "price-range",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filters restaurants based on a list districts, separated by commas. A restaurant will be returned if it is in one of the provided districts. Example: district=Praha 1,Praha2",
                        "name": "district",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all non vegetarian restaurants.",
                        "name": "vegetarian",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all non vegan restaurants.",
                        "name": "vegan",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all non gluten free restaurants.",
                        "name": "gluten-free",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all restaurants that don't have a takeaway option.",
                        "name": "takeaway",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "Filters out all restaurants that don't have a delivery option.",
                        "name": "delivery-options",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sorts restaurants. Available sort options: price-asc, price-desc, rating",
                        "name": "sort",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseFullJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Returns a JSON with user info if the request headers contain an authenticated cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get info about a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseUserJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a user if the request headers contain an authenticated cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Deletes a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates user's password or username based on the provided JSON. Only 1 field can be updated at a time. For password you need to provide \"oldPassword\" and \"newPassword\" fields, omitting the \"newUsername\" field and vice versa if you'd like to update the username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Updates a user's password or username",
                "parameters": [
                    {
                        "description": "Create a new user",
                        "name": "updateJSON",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.userUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.responseSimpleJSON"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.responseAutocompleteJSON": {
            "type": "object",
            "properties": {
                "Data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.restaurantAutocomplete"
                    }
                },
                "Msg": {
                    "type": "string",
                    "example": "Success"
                },
                "Status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "api.responseFullJSON": {
            "type": "object",
            "properties": {
                "Data": {
                    "type": "string"
                },
                "Msg": {
                    "type": "string",
                    "example": "Success"
                },
                "Status": {
                    "type": "integer",
                    "example": 200
                },
                "User": {
                    "type": "object",
                    "$ref": "#/definitions/api.userResponse"
                }
            }
        },
        "api.responseSimpleJSON": {
            "type": "object",
            "properties": {
                "Msg": {
                    "type": "string"
                },
                "Status": {
                    "type": "integer"
                }
            }
        },
        "api.responseUserJSON": {
            "type": "object",
            "properties": {
                "Msg": {
                    "type": "string",
                    "example": "Success"
                },
                "Status": {
                    "type": "integer",
                    "example": 200
                },
                "User": {
                    "type": "object",
                    "$ref": "#/definitions/api.userResponse"
                }
            }
        },
        "api.restaurantAutocomplete": {
            "type": "object",
            "properties": {
                "Address": {
                    "type": "string",
                    "example": "Polská 13"
                },
                "District": {
                    "type": "string",
                    "example": "Praha 1"
                },
                "ID": {
                    "type": "integer",
                    "example": 1
                },
                "Image": {
                    "type": "string",
                    "example": "url.com"
                },
                "Name": {
                    "type": "string",
                    "example": "Steakhouse"
                }
            }
        },
        "api.userResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "api.userUpdate": {
            "type": "object",
            "required": [
                "newPassword",
                "newUsername",
                "oldPassword"
            ],
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "newUsername": {
                    "type": "string"
                },
                "oldPassword": {
                    "type": "string"
                }
            }
        },
        "db.User": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
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
	Version:     "0.2.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Restaurateur API",
	Description: "Provides info about restaurants in Prague",
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
