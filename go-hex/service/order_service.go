package service

import (
	"errors"
	"hex/model"
	"hex/repo"
)

type OrderService interface {
	CreateOrder(order model.Order) error
}

type orderServiceImpl struct {
	repo repo.OrderRepository
}

func NewOrderService(repo repo.OrderRepository) OrderService {
	return &orderServiceImpl{repo: repo}
}

func (s *orderServiceImpl) CreateOrder(order model.Order) error {
	if order.Total <= 0 {
		return errors.New("Total must be greater than 0")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}
