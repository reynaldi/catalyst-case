package customer

import (
	"catalyst-case/database"
	"context"
	"errors"
)

const (
	addNewCustomer = `INSERT INTO customers (email, created_at) VALUES (?, ?)`
)

type customerCommand struct {
	*database.DB
}

type CustomerCommand interface {
	AddCustomer(ctx context.Context, customer CustomerEntity) error
}

func NewCustomerCommand(db *database.DB) CustomerCommand {
	return &customerCommand{
		DB: db,
	}
}

func (c *customerCommand) AddCustomer(ctx context.Context, customer CustomerEntity) error {
	res, e := c.ExecContext(ctx, addNewCustomer, customer.Email, customer.CreatedAt)
	if e != nil {
		return e
	}
	affected, e := res.RowsAffected()
	if e != nil {
		return e
	}
	if affected == 0 {
		return errors.New("failed to add new customer")
	}
	return nil
}
