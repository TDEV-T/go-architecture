package adapter

import (
	"clean/entities"
	"clean/repositories"

	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) repositories.ProductRepo {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) Save(order entities.Product) error {
	return r.db.Create(&order).Error
}
