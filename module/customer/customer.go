package customer

import (
	"context"
	"time"
)

type customerService struct {
	query   CustomerQuery
	command CustomerCommand
}

type Customer interface {
	AddNewCustomer(ctx context.Context, customer NewCustomerDto) error
	GetCustomerById(ctx context.Context, customerId int) (*CustomerDto, error)
	GetCustomerByEmail(ctx context.Context, email string) (*CustomerDto, error)
}

func NewCustomer(query CustomerQuery, command CustomerCommand) Customer {
	return &customerService{
		query:   query,
		command: command,
	}
}

func (c *customerService) AddNewCustomer(ctx context.Context, customer NewCustomerDto) error {
	var entity = mapNewCustomerToEntity(customer)
	var e = c.command.AddCustomer(ctx, entity)
	return e
}

func (c *customerService) GetCustomerById(ctx context.Context, customerId int) (*CustomerDto, error) {
	customer, e := c.query.GetCustomerById(ctx, customerId)
	if e != nil {
		return nil, e
	}
	if customer == nil {
		return nil, nil
	}
	var dto = mapCustomerEntityToDto(customer)
	return dto, nil
}

func (c *customerService) GetCustomerByEmail(ctx context.Context, email string) (*CustomerDto, error) {
	customer, e := c.query.GetCustomerByEmail(ctx, email)
	if e != nil {
		return nil, e
	}
	if customer == nil {
		return nil, nil
	}
	var dto = mapCustomerEntityToDto(customer)
	return dto, nil
}

func mapNewCustomerToEntity(cust NewCustomerDto) CustomerEntity {
	var result = CustomerEntity{
		Email:     cust.Email,
		Name:      cust.Name,
		CreatedAt: time.Now().UTC(),
	}
	return result
}

func mapCustomerEntityToDto(customer *CustomerEntity) *CustomerDto {
	var result = &CustomerDto{
		CustomerId: customer.CustomerId,
		Email:      customer.Email,
		Name:       customer.Name,
	}
	return result
}
