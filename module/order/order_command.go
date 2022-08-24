package order

import (
	"catalyst-case/database"
	"context"
	"log"
)

const (
	addNewMaster = `INSERT INTO orders (grand_total, created_at, created_by) VALUES (?, ?, ?)`
	addNewDetail = `INSERT INTO order_details (order_id, product_id, product_name, unit_price, qty) VALUES (?, ?, ?, ?, ?)`
)

type orderCommand struct {
	*database.DB
}

type OrderCommand interface {
	AddNewOrder(ctx context.Context, request OrderMasterEntity, detail []OrderDetailEntity) (*OrderWrap, error)
}

func NewOrderCommand(db *database.DB) OrderCommand {
	return &orderCommand{
		DB: db,
	}
}

func (o *orderCommand) AddNewOrder(ctx context.Context, request OrderMasterEntity, detail []OrderDetailEntity) (*OrderWrap, error) {
	tx, e := o.BeginTx(ctx, nil)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	defer tx.Rollback()
	// Create master
	masterResult, e := tx.ExecContext(ctx, addNewMaster, request.GrandTotal, request.CreatedAt, request.CreatedBy)
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	masterId, e := masterResult.LastInsertId()
	if e != nil {
		tx.Rollback()
		return nil, e
	}
	request.OrderId = int(masterId)
	// Iterate insert details
	for idx, item := range detail {
		detResult, e := tx.ExecContext(ctx, addNewDetail, request.OrderId, item.ProductId, item.ProductName, item.UnitPrice, item.Qty)
		if e != nil {
			tx.Rollback()
			return nil, e
		}
		detId, e := detResult.LastInsertId()
		if e != nil {
			tx.Rollback()
			return nil, e
		}
		detail[idx].OrderDetailId = int(detId)
	}
	err := tx.Commit()
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}
	wrap := &OrderWrap{
		Order:        request,
		OrderDetails: detail,
	}
	return wrap, nil
}
