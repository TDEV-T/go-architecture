package adapter

import (
	"hex/model"
	"hex/service"

	"github.com/gofiber/fiber/v2"
)

type HttpOrderHandler struct {
	service service.OrderService
}

func NewHttpOrderHandler(service service.OrderService) *HttpOrderHandler {
	return &HttpOrderHandler{service: service}
}

func (h *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	var order model.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Order created successfully"})

}
