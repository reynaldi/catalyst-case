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

type OrderWrap struct {
	Order        OrderMasterEntity
	OrderDetails []OrderDetailEntity
}

type OrderEntry struct {
	Items     []OrderItem `json:"items"`
	CreatedBy int         `json:"created_by"`
}

type OrderItem struct {
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

type OrderDto struct {
	OrderId    int              `json:"order_id"`
	GrandTotal float64          `json:"grand_total"`
	CreatedAt  time.Time        `json:"created_at"`
	Items      []OrderDetailDto `json:"items"`
	CreatedBy  int              `json:"created_by"`
}

type OrderDetailDto struct {
	OrderDetailId int     `json:"order_detail_id"`
	OrderId       int     `json:"order_id"`
	ProductName   string  `json:"product_name"`
	Qty           int     `json:"qty"`
	UnitPrice     float64 `json:"unit_price"`
}
