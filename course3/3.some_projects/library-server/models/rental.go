package models

type Rental interface{}

type ConcreteRental struct {
	ID         int          `json:"id"`
	Book       ConcreteBook `json:"book"`
	User       ConcreteUser `json:"user"`
	RentedAt   string       `json:"rented_at"`
	ReturnedAt *string      `json:"returned_at"`
}
