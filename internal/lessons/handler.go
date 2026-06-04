package lessons

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(c *fiber.Ctx) error {
	lessons, err := h.service.List()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    lessons,
	})
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	lesson, err := h.service.FindByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "lesson not found",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"data":    lesson,
	})
}
