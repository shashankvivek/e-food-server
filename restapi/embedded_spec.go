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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "E-Food",
    "version": "1.0.0"
  },
  "host": "e-food.com",
  "basePath": "/v1",
  "paths": {
    "/categories": {
      "get": {
        "tags": [
          "menu"
        ],
        "operationId": "CategoryList",
        "responses": {
          "200": {
            "description": "Get Category to show menu",
            "schema": {
              "$ref": "#/definitions/Categories"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Categories Not Found"
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/checkoutCart": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get Checkout Cart with offers",
        "tags": [
          "user"
        ],
        "operationId": "checkout",
        "responses": {
          "200": {
            "description": "Success response when item is added successfully",
            "schema": {
              "$ref": "#/definitions/BillableCart"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be added Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/guest/cart": {
      "get": {
        "description": "Get All cart items",
        "tags": [
          "guest"
        ],
        "operationId": "GetItems",
        "responses": {
          "200": {
            "description": "All items in cart",
            "schema": {
              "$ref": "#/definitions/CartPreview"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "description": "This API adds product to cart / guest cart",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "guest"
        ],
        "operationId": "AddItem",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ItemInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is added successfully",
            "schema": {
              "$ref": "#/definitions/CartSuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be added Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "description": "Remove item from cart",
        "tags": [
          "guest"
        ],
        "operationId": "RemoveItem",
        "parameters": [
          {
            "type": "integer",
            "name": "productId",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is removed from cart",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be deleted Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "Returns token for authorized User",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Login",
        "parameters": [
          {
            "description": "Login Payload",
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful login",
            "schema": {
              "$ref": "#/definitions/LoginSuccess"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/productListBySubCategory/{id}": {
      "get": {
        "tags": [
          "products"
        ],
        "operationId": "GetFromSubCategory",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Subcategory Id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Get Products based on sub category",
            "schema": {
              "$ref": "#/definitions/Products"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Products not found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "description": "To register a new user",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Register",
        "parameters": [
          {
            "description": "Registeration Payload",
            "name": "signup",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful registeration",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/sessionInfo": {
      "post": {
        "description": "Adds Cookie ID for guests",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "guest"
        ],
        "operationId": "AddSession",
        "parameters": [
          {
            "name": "session_info",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/GuestSession"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is added successfully",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Session ID Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/user/cart": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get All cart items",
        "tags": [
          "user"
        ],
        "operationId": "GetCart",
        "responses": {
          "200": {
            "description": "All items in cart",
            "schema": {
              "$ref": "#/definitions/CartPreview"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "This API adds product to cart / guest cart",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "AddToCart",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ItemInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is added successfully",
            "schema": {
              "$ref": "#/definitions/CartSuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be added Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Remove item from cart",
        "tags": [
          "user"
        ],
        "operationId": "RemoveFromCart",
        "parameters": [
          {
            "type": "integer",
            "name": "productId",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is removed from cart",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be deleted Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BillableCart": {
      "type": "object",
      "properties": {
        "currency": {
          "type": "string"
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BillingItem"
          }
        },
        "offerItems": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/OfferItem"
          }
        },
        "totalPrice": {
          "type": "number"
        },
        "totalSaving": {
          "type": "number"
        }
      }
    },
    "BillingItem": {
      "type": "object",
      "properties": {
        "currency": {
          "type": "string"
        },
        "imageUrl": {
          "type": "string"
        },
        "productId": {
          "type": "integer"
        },
        "productName": {
          "type": "string"
        },
        "quantity": {
          "type": "integer"
        },
        "totalPrice": {
          "type": "number"
        },
        "unitPrice": {
          "type": "number"
        }
      }
    },
    "CartItem": {
      "type": "object",
      "properties": {
        "currency": {
          "type": "string"
        },
        "imageUrl": {
          "type": "string"
        },
        "productId": {
          "type": "integer"
        },
        "productName": {
          "type": "string"
        },
        "quantity": {
          "type": "integer"
        },
        "unitPrice": {
          "type": "number"
        }
      }
    },
    "CartPreview": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/CartItem"
      }
    },
    "CartSuccessResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "qtyAdded": {
          "type": "integer"
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "Categories": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Category"
      }
    },
    "Category": {
      "type": "object",
      "properties": {
        "bcId": {
          "type": "integer"
        },
        "bcImageUrl": {
          "type": "string"
        },
        "bcIsActive": {
          "type": "boolean"
        },
        "bcName": {
          "type": "string"
        },
        "subCategories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/SubCategory"
          }
        }
      }
    },
    "GuestSession": {
      "type": "object",
      "properties": {
        "extraInfo": {
          "type": "string"
        }
      }
    },
    "ItemInfo": {
      "type": "object",
      "properties": {
        "productId": {
          "type": "integer"
        },
        "totalQty": {
          "type": "integer"
        }
      }
    },
    "LoginInfo": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "OfferItem": {
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BillingItem"
          }
        },
        "ruleSetId": {
          "type": "string"
        }
      }
    },
    "Product": {
      "type": "object",
      "properties": {
        "bcId": {
          "description": "Broad Category Id",
          "type": "integer"
        },
        "currency": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "discountPercentage": {
          "description": "Discount to be applied on Unit Price",
          "type": "number",
          "example": "1.00"
        },
        "imageUrl": {
          "type": "string"
        },
        "isAvailable": {
          "description": "False if Product is out of stock",
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "productId": {
          "type": "integer"
        },
        "scId": {
          "description": "Sub Category Id",
          "type": "integer"
        },
        "sku": {
          "type": "string"
        },
        "unitPrice": {
          "type": "number"
        }
      }
    },
    "Products": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Product"
      }
    },
    "RegisterUser": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "fname": {
          "type": "string"
        },
        "lname": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phoneNo": {
          "type": "string"
        }
      }
    },
    "SubCategory": {
      "type": "object",
      "properties": {
        "scId": {
          "type": "integer"
        },
        "scImageUrl": {
          "type": "string"
        },
        "scIsActive": {
          "type": "boolean"
        },
        "scName": {
          "type": "string"
        }
      }
    },
    "SuccessResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "E-Food",
    "version": "1.0.0"
  },
  "host": "e-food.com",
  "basePath": "/v1",
  "paths": {
    "/categories": {
      "get": {
        "tags": [
          "menu"
        ],
        "operationId": "CategoryList",
        "responses": {
          "200": {
            "description": "Get Category to show menu",
            "schema": {
              "$ref": "#/definitions/Categories"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Categories Not Found"
          },
          "500": {
            "description": "Server Error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/checkoutCart": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get Checkout Cart with offers",
        "tags": [
          "user"
        ],
        "operationId": "checkout",
        "responses": {
          "200": {
            "description": "Success response when item is added successfully",
            "schema": {
              "$ref": "#/definitions/BillableCart"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be added Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/guest/cart": {
      "get": {
        "description": "Get All cart items",
        "tags": [
          "guest"
        ],
        "operationId": "GetItems",
        "responses": {
          "200": {
            "description": "All items in cart",
            "schema": {
              "$ref": "#/definitions/CartPreview"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "description": "This API adds product to cart / guest cart",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "guest"
        ],
        "operationId": "AddItem",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ItemInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is added successfully",
            "schema": {
              "$ref": "#/definitions/CartSuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be added Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "description": "Remove item from cart",
        "tags": [
          "guest"
        ],
        "operationId": "RemoveItem",
        "parameters": [
          {
            "type": "integer",
            "name": "productId",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is removed from cart",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be deleted Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "Returns token for authorized User",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Login",
        "parameters": [
          {
            "description": "Login Payload",
            "name": "login",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/LoginInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful login",
            "schema": {
              "$ref": "#/definitions/LoginSuccess"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/productListBySubCategory/{id}": {
      "get": {
        "tags": [
          "products"
        ],
        "operationId": "GetFromSubCategory",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "description": "Subcategory Id",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Get Products based on sub category",
            "schema": {
              "$ref": "#/definitions/Products"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Products not found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/register": {
      "post": {
        "description": "To register a new user",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "Register",
        "parameters": [
          {
            "description": "Registeration Payload",
            "name": "signup",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/RegisterUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful registeration",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "User not found",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/sessionInfo": {
      "post": {
        "description": "Adds Cookie ID for guests",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "guest"
        ],
        "operationId": "AddSession",
        "parameters": [
          {
            "name": "session_info",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/GuestSession"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is added successfully",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Session ID Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/user/cart": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get All cart items",
        "tags": [
          "user"
        ],
        "operationId": "GetCart",
        "responses": {
          "200": {
            "description": "All items in cart",
            "schema": {
              "$ref": "#/definitions/CartPreview"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "This API adds product to cart / guest cart",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "operationId": "AddToCart",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/ItemInfo"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is added successfully",
            "schema": {
              "$ref": "#/definitions/CartSuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be added Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Remove item from cart",
        "tags": [
          "user"
        ],
        "operationId": "RemoveFromCart",
        "parameters": [
          {
            "type": "integer",
            "name": "productId",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success response when item is removed from cart",
            "schema": {
              "$ref": "#/definitions/SuccessResponse"
            }
          },
          "400": {
            "description": "Bad Request"
          },
          "404": {
            "description": "Item to be deleted Not Found"
          },
          "500": {
            "description": "Server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "BillableCart": {
      "type": "object",
      "properties": {
        "currency": {
          "type": "string"
        },
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BillingItem"
          }
        },
        "offerItems": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/OfferItem"
          }
        },
        "totalPrice": {
          "type": "number"
        },
        "totalSaving": {
          "type": "number"
        }
      }
    },
    "BillingItem": {
      "type": "object",
      "properties": {
        "currency": {
          "type": "string"
        },
        "imageUrl": {
          "type": "string"
        },
        "productId": {
          "type": "integer"
        },
        "productName": {
          "type": "string"
        },
        "quantity": {
          "type": "integer"
        },
        "totalPrice": {
          "type": "number"
        },
        "unitPrice": {
          "type": "number"
        }
      }
    },
    "CartItem": {
      "type": "object",
      "properties": {
        "currency": {
          "type": "string"
        },
        "imageUrl": {
          "type": "string"
        },
        "productId": {
          "type": "integer"
        },
        "productName": {
          "type": "string"
        },
        "quantity": {
          "type": "integer"
        },
        "unitPrice": {
          "type": "number"
        }
      }
    },
    "CartPreview": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/CartItem"
      }
    },
    "CartSuccessResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "qtyAdded": {
          "type": "integer"
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "Categories": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Category"
      }
    },
    "Category": {
      "type": "object",
      "properties": {
        "bcId": {
          "type": "integer"
        },
        "bcImageUrl": {
          "type": "string"
        },
        "bcIsActive": {
          "type": "boolean"
        },
        "bcName": {
          "type": "string"
        },
        "subCategories": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/SubCategory"
          }
        }
      }
    },
    "GuestSession": {
      "type": "object",
      "properties": {
        "extraInfo": {
          "type": "string"
        }
      }
    },
    "ItemInfo": {
      "type": "object",
      "properties": {
        "productId": {
          "type": "integer"
        },
        "totalQty": {
          "type": "integer"
        }
      }
    },
    "LoginInfo": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "LoginSuccess": {
      "type": "object",
      "properties": {
        "success": {
          "type": "boolean"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "OfferItem": {
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/BillingItem"
          }
        },
        "ruleSetId": {
          "type": "string"
        }
      }
    },
    "Product": {
      "type": "object",
      "properties": {
        "bcId": {
          "description": "Broad Category Id",
          "type": "integer"
        },
        "currency": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "discountPercentage": {
          "description": "Discount to be applied on Unit Price",
          "type": "number",
          "example": "1.00"
        },
        "imageUrl": {
          "type": "string"
        },
        "isAvailable": {
          "description": "False if Product is out of stock",
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "productId": {
          "type": "integer"
        },
        "scId": {
          "description": "Sub Category Id",
          "type": "integer"
        },
        "sku": {
          "type": "string"
        },
        "unitPrice": {
          "type": "number"
        }
      }
    },
    "Products": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Product"
      }
    },
    "RegisterUser": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "fname": {
          "type": "string"
        },
        "lname": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "phoneNo": {
          "type": "string"
        }
      }
    },
    "SubCategory": {
      "type": "object",
      "properties": {
        "scId": {
          "type": "integer"
        },
        "scImageUrl": {
          "type": "string"
        },
        "scIsActive": {
          "type": "boolean"
        },
        "scName": {
          "type": "string"
        }
      }
    },
    "SuccessResponse": {
      "type": "object",
      "properties": {
        "message": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
