package order

import "catalyst-case/database"

type orderCommand struct {
	*database.DB
}

type OrderCommand interface {
}
