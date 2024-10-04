package usecases

import (
	"clean/entities"
	"clean/repositories"
)

type ProductUseCase interface {
	CreateProduct(order entities.Product) error
}

type ProductService struct {
	repo repositories.ProductRepo
}

func NewOrderService(repo repositories.ProductRepo) ProductUseCase {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(order entities.Product) error {
	err := s.repo.Save(order)
	if err != nil {
		return err
	}
	return nil
}
