// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Secret Santa",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/group": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "groups"
        ],
        "summary": "Add a new group",
        "parameters": [
          {
            "description": "Group Object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Group"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Great"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/group/{id}": {
      "get": {
        "tags": [
          "groups"
        ],
        "summary": "Gets a groups info",
        "parameters": [
          {
            "type": "string",
            "description": "Group Id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Group"
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    },
    "/member": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "members"
        ],
        "summary": "Adding a member with a group",
        "parameters": [
          {
            "name": "member",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Member"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Member"
            }
          },
          "400": {
            "description": "Bad Request"
          }
        }
      }
    }
  },
  "definitions": {
    "Group": {
      "type": "object",
      "required": [
        "id",
        "deadline",
        "exchangeDate",
        "moneyLimit"
      ],
      "properties": {
        "deadline": {
          "type": "string"
        },
        "exchangeDate": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "moneyLimit": {
          "type": "number"
        }
      }
    },
    "Member": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "group": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "spouse": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Secret Santa",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/group": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "groups"
        ],
        "summary": "Add a new group",
        "parameters": [
          {
            "description": "Group Object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Group"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Great"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/group/{id}": {
      "get": {
        "tags": [
          "groups"
        ],
        "summary": "Gets a groups info",
        "parameters": [
          {
            "type": "string",
            "description": "Group Id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Group"
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    },
    "/member": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "members"
        ],
        "summary": "Adding a member with a group",
        "parameters": [
          {
            "name": "member",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Member"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Member"
            }
          },
          "400": {
            "description": "Bad Request"
          }
        }
      }
    }
  },
  "definitions": {
    "Group": {
      "type": "object",
      "required": [
        "id",
        "deadline",
        "exchangeDate",
        "moneyLimit"
      ],
      "properties": {
        "deadline": {
          "type": "string"
        },
        "exchangeDate": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "moneyLimit": {
          "type": "number"
        }
      }
    },
    "Member": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "group": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "spouse": {
          "type": "string"
        }
      }
    }
  }
}`))
}
