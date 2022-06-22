package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"` //validator package
	Price       json.Number `json:"price" binding:"required,number"`
	Rating      json.Number `json:"rating" binding:"number"` //validator package
	Description string      `json:"description"`             //validator package
}
