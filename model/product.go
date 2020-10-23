package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// Product struct represents product model
type Product struct {
	Id          int    `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Category    string `json: "category"`
	Amount      int    `json: "amount"`
}

// Products struct represents products model
type Products struct {
	Products []Product `json: "products"`
}

// ValidateRequest func to validate a request
func (p Product) ValidateRequest() error {
	return validation.ValidateStruct(&p,
		// Name is required, and the length must between 5 and 50
		validation.Field(&p.Name, validation.Required, validation.Length(5, 50)),
		// Description is required, and the minimum length is 10
		validation.Field(&p.Description, validation.Required, validation.Length(10, 0)),
		// Category is required
		validation.Field(&p.Category, validation.Required),
		// Amount is required, and the minimum value is equals 1
		validation.Field(&p.Amount, validation.Required, validation.Min(1)),
	)
}
