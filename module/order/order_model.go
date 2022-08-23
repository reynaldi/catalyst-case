package order

import (
	"database/sql"
	"time"
)

type OrderMasterEntity struct {
	OrderId    int             `db:"order_id"`
	GrandTotal sql.NullFloat64 `db:"grand_total,omitempty"`
	CreatedAt  time.Time       `db:"created_at"`
	CreatedBy  int             `db:"created_by"`
}

type OrderDetailEntity struct {
	OrderDetailId int     `db:"order_detail_id"`
	OrderId       int     `db:"order_id"`
	ProductId     int     `db:"product_id"`
	ProductName   string  `db:"product_name"`
	UnitPrice     float64 `db:"unit_price"`
	Qty           int     `db:"qty"`
}

type OrderRequest struct {
	Items []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}
