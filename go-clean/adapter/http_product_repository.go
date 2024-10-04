package adapter

import (
	"clean/entities"
	"clean/usecases"

	"github.com/gofiber/fiber/v2"
)

type HttpProductHandler struct {
	productUseCase usecases.ProductUseCase
}

func NewHttpProductHandler(productUseCase usecases.ProductUseCase) *HttpProductHandler {
	return &HttpProductHandler{productUseCase: productUseCase}
}

func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	var order entities.Product
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.productUseCase.CreateProduct(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
