// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/barber/barberShop/create": {
            "post": {
                "description": "Cria um estavelecimento",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barberShop"
                ],
                "summary": "Criação das barbearias",
                "parameters": [
                    {
                        "description": "Create barber shop",
                        "name": "barberShop",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/barberShop.CreateBarberShop"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sem conteúdo"
                    }
                }
            }
        },
        "/barber/barberShop/list": {
            "get": {
                "description": "Lista todas as barbearias disponíveis para serem feiras as reservas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barberShop"
                ],
                "summary": "Lista todas as barbearias",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/barberShop.ListBarbserShop"
                            }
                        }
                    }
                }
            }
        },
        "/barber/client/create": {
            "post": {
                "description": "Cria um novo cliente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "Create client",
                "parameters": [
                    {
                        "description": "Create client",
                        "name": "client",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/client.CreateClient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sem conteúdo"
                    }
                }
            }
        },
        "/barber/client/list": {
            "get": {
                "description": "Lista todos os clientes da barbearia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "client"
                ],
                "summary": "List os clientes da barbearia",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/client.ListClients"
                            }
                        }
                    }
                }
            }
        },
        "/barber/create": {
            "post": {
                "description": "Cria um novo barbeiro",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Criação dos barbeiros",
                "parameters": [
                    {
                        "description": "Create barber",
                        "name": "barber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/barber.CreateBarber"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sem conteúdo"
                    }
                }
            }
        },
        "/barber/list": {
            "get": {
                "description": "Lista todos os barbeiros da barbearia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "barber"
                ],
                "summary": "Lista os barbeiros da barbearia",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/barber.ListBarbers"
                            }
                        }
                    }
                }
            }
        },
        "/barber/service/create": {
            "post": {
                "description": "Cria um novo serviço para a barbeiro e barbearia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Criação dos serviços",
                "parameters": [
                    {
                        "description": "Create service",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.CreateService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Sem conteúdo"
                    }
                }
            }
        },
        "/barber/service/list": {
            "get": {
                "description": "Lista todos os serviços ofertados pela barbearia",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "service"
                ],
                "summary": "Lista os serviços da barbearia",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/service.ListServices"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "barber.CreateBarber": {
            "type": "object",
            "required": [
                "contato",
                "nome"
            ],
            "properties": {
                "contato": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                }
            }
        },
        "barber.ListBarbers": {
            "type": "object",
            "properties": {
                "contato": {
                    "type": "string"
                },
                "data_atualizacao": {
                    "type": "string"
                },
                "data_criacao": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                }
            }
        },
        "barberShop.CreateBarberShop": {
            "type": "object",
            "required": [
                "cidade",
                "contato",
                "nome",
                "numero_residencia",
                "ponto_referencia",
                "rua"
            ],
            "properties": {
                "cidade": {
                    "type": "string"
                },
                "contato": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "numero_residencia": {
                    "type": "integer"
                },
                "ponto_referencia": {
                    "type": "string"
                },
                "rua": {
                    "type": "string"
                }
            }
        },
        "barberShop.ListBarbserShop": {
            "type": "object",
            "properties": {
                "cidade": {
                    "type": "string"
                },
                "contato": {
                    "type": "string"
                },
                "data_atualizacao": {
                    "type": "string"
                },
                "data_criacao": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "numero_residencia": {
                    "type": "integer"
                },
                "ponto_referencia": {
                    "type": "string"
                },
                "rua": {
                    "type": "string"
                }
            }
        },
        "client.CreateClient": {
            "type": "object",
            "required": [
                "contato",
                "email",
                "nome",
                "senha"
            ],
            "properties": {
                "contato": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "senha": {
                    "type": "string"
                }
            }
        },
        "client.ListClients": {
            "type": "object",
            "properties": {
                "contato": {
                    "type": "string"
                },
                "data_atualizacao": {
                    "type": "string"
                },
                "data_criacao": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "senha": {
                    "type": "string"
                }
            }
        },
        "service.CreateService": {
            "type": "object",
            "required": [
                "nome",
                "preco"
            ],
            "properties": {
                "duracao": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "preco": {
                    "type": "number"
                }
            }
        },
        "service.ListServices": {
            "type": "object",
            "properties": {
                "data_atualizacao": {
                    "type": "string"
                },
                "data_criacao": {
                    "type": "string"
                },
                "duracao": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "preco": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Barber Shop API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	//LeftDelim:        "{{",
	//RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
