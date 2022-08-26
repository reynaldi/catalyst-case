package customer

import (
	"catalyst-case/database"
	"context"
	"database/sql"
)

const (
	getCustomerById    = `SELECT customer_id, email, name, created_at FROM customers WHERE customer_id = ?`
	getCustomerByEmail = `SELECT customer_id, email, name, created_at FROM customers WHERE email = ?`
)

type customerQuery struct {
	*database.DB
}

type CustomerQuery interface {
	GetCustomerById(ctx context.Context, customerId int) (*CustomerEntity, error)
	GetCustomerByEmail(ctx context.Context, email string) (*CustomerEntity, error)
}

func NewCustomerQuery(db *database.DB) CustomerQuery {
	return &customerQuery{
		DB: db,
	}
}

func (c *customerQuery) GetCustomerById(ctx context.Context, customerId int) (*CustomerEntity, error) {
	row := c.QueryRowContext(ctx, getCustomerById, customerId)
	result, e := ScanCustomer(row)
	if e != nil {
		return nil, e
	}
	return result, nil
}

func (c *customerQuery) GetCustomerByEmail(ctx context.Context, email string) (*CustomerEntity, error) {
	row := c.QueryRowContext(ctx, getCustomerByEmail, email)
	result, e := ScanCustomer(row)
	if e != nil {
		return nil, e
	}
	return result, nil
}

func ScanCustomer(row *sql.Row) (*CustomerEntity, error) {
	var result CustomerEntity
	e := row.Scan(&result.CustomerId, &result.Email, &result.Name, &result.CreatedAt)
	if e != nil {
		if e == sql.ErrNoRows {
			return nil, nil
		}
		return nil, e
	}
	return &result, nil
}
