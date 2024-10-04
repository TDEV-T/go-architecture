package adapter

import (
	"hex/model"

	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) *GormOrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order model.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}

	return nil
}
