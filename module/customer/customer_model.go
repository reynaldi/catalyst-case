package customer

import "time"

type CustomerEntity struct {
	CustomerId int       `db:"customer_id"`
	Email      string    `db:"email"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
}

type NewCustomerDto struct {
	Email string `json:"email" validate:"required"`
	Name  string `json:"name"`
}

type CustomerDto struct {
	CustomerId int    `json:"customer_id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
}
