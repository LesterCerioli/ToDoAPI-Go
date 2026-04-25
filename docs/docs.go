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
            "name": "API Support",
            "url": "https://github.com/yourusername/todo-api",
            "email": "support@example.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/todos": {
            "get": {
                "description": "Retrieve all todo items from database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "List all todos",
                "responses": {
                    "200": {
                        "description": "List of todos",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new todo item with title and optional description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Create a new todo",
                "parameters": [
                    {
                        "description": "Todo creation payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created todo",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Retrieve a specific todo item by its UUID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Get a todo by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Todo item",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Partially update a todo item (title, description, or completed status)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Update a todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated todo",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Soft-delete a todo item (or hard-delete depending on config)",
                "tags": [
                    "todos"
                ],
                "summary": "Delete a todo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateTodoRequest": {
            "description": "Request body for creating a todo",
            "type": "object",
            "properties": {
                "description": {
                    "description": "Todo description",
                    "type": "string",
                    "example": "Build a CRUD app with Swagger"
                },
                "title": {
                    "description": "Todo title",
                    "type": "string",
                    "example": "Learn Fiber v3"
                }
            }
        },
        "models.UpdateTodoRequest": {
            "description": "Request body for partial updates (all fields optional)",
            "type": "object",
            "properties": {
                "completed": {
                    "description": "Completion status",
                    "type": "boolean",
                    "example": true
                },
                "description": {
                    "description": "Updated description",
                    "type": "string",
                    "example": "Updated description"
                },
                "title": {
                    "description": "Updated title",
                    "type": "string",
                    "example": "Updated title"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "todoapi-go.onrender.com",
	BasePath:         "/api/v1",
	Schemes:          []string{"https"},
	Title:            "Todo API",
	Description:      "A production-ready CRUD Todo API with SQLite, Swagger UI, and graceful shutdown.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
