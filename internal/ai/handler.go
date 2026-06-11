package ai

import (
	"net/http"

	"akubisa/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Handler handles HTTP requests for AI analysis
type Handler interface {
	AnalyzePronunciation(c *fiber.Ctx) error
}

type handler struct {
	service Service
}

// NewHandler creates a new AI handler
func NewHandler(service Service) Handler {
	return &handler{service: service}
}

// AnalyzePronunciation handles the POST /api/v1/ai/pronunciation/analyze request
func (h *handler) AnalyzePronunciation(c *fiber.Ctx) error {
	childIDStr := c.FormValue("child_id")
	expectedText := c.FormValue("expected_text")

	if childIDStr == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "child_id is required",
		})
	}

	childID, err := uuid.Parse(childIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid child_id format",
		})
	}

	fileHeader, err := c.FormFile("audio")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "audio file is required",
			"error":   err.Error(),
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "failed to open audio file",
		})
	}
	defer file.Close()

	result, err := h.service.AnalyzePronunciation(c.Context(), childID, file, fileHeader.Filename, expectedText)
	if err != nil {
		logger.Logger.Error("failed to analyze pronunciation", "error", err, "child_id", childIDStr)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "failed to analyze pronunciation",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Pronunciation analyzed successfully",
		"data":    result,
	})
}
