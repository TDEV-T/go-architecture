package repo

import "hex/model"

type OrderRepository interface {
	Save(order model.Order) error
}
