package repositories

import "clean/entities"

type ProductRepo interface {
	Save(order entities.Product) error
}
