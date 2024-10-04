package route

import (
	"hex/adapter"
	"hex/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitRoute(route *fiber.App, db *gorm.DB) {
	orderRepo := adapter.NewGormOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := adapter.NewHttpOrderHandler(orderService)

	route.Post("/order", orderHandler.CreateOrder)
}
