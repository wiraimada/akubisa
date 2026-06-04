package lessons

import (
	"net/http"

	"akubisa/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Handler handles HTTP requests for lessons
type Handler interface {
	ListLessons(c *fiber.Ctx) error
	GetLesson(c *fiber.Ctx) error
	SubmitQuiz(c *fiber.Ctx) error
}

type handler struct {
	service Service
}

// NewHandler creates a new lesson handler
func NewHandler(service Service) Handler {
	return &handler{service: service}
}

// ListLessons godoc
// @Summary List all lessons
// @Description Get a list of all available lessons with their quizzes and questions
// @Tags lessons
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "List of lessons"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/lessons [get]
func (h *handler) ListLessons(c *fiber.Ctx) error {
	lessons, err := h.service.ListLessons()
	if err != nil {
		logger.Logger.Error("failed to list lessons", "error", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve lessons",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Lessons retrieved successfully",
		"data":    lessons,
	})
}

// GetLesson godoc
// @Summary Get lesson details by ID
// @Description Get details of a specific lesson, including its quizzes and questions
// @Tags lessons
// @Accept json
// @Produce json
// @Param id path string true "Lesson ID"
// @Success 200 {object} map[string]interface{} "Lesson details"
// @Failure 400 {object} map[string]interface{} "Invalid Lesson ID"
// @Failure 404 {object} map[string]interface{} "Lesson not found"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/lessons/{id} [get]
func (h *handler) GetLesson(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid lesson ID format",
			"error":   err.Error(),
		})
	}

	lesson, err := h.service.GetLesson(id)
	if err != nil {
		if err.Error() == "lesson not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Lesson not found",
			})
		}
		logger.Logger.Error("failed to get lesson", "error", err, "lesson_id", idStr)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to retrieve lesson",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Lesson retrieved successfully",
		"data":    lesson,
	})
}

// SubmitQuizRequest represents the request body for submitting a quiz
type SubmitQuizRequest struct {
	ChildID     uuid.UUID        `json:"child_id" validate:"required"`
	Submissions []QuizSubmission `json:"submissions" validate:"required,min=1"`
}

// SubmitQuiz godoc
// @Summary Submit answers for a quiz
// @Description Submit a child's answers to a specific quiz and record their progress
// @Tags quizzes
// @Accept json
// @Produce json
// @Param id path string true "Quiz ID"
// @Param request body SubmitQuizRequest true "Quiz submission data"
// @Success 200 {object} map[string]interface{} "Quiz submitted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request or quiz ID"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/v1/quizzes/{id}/submit [post]
func (h *handler) SubmitQuiz(c *fiber.Ctx) error {
	quizIDStr := c.Params("id")
	quizID, err := uuid.Parse(quizIDStr)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid quiz ID format",
			"error":   err.Error(),
		})
	}

	var req SubmitQuizRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Basic validation for ChildID
	if req.ChildID == uuid.Nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Child ID is required",
		})
	}

	progress, err := h.service.SubmitQuiz(quizID, req.ChildID, req.Submissions)
	if err != nil {
		logger.Logger.Error("failed to submit quiz", "error", err, "quiz_id", quizIDStr, "child_id", req.ChildID)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to submit quiz",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Quiz submitted and progress recorded successfully",
		"data":    progress,
	})
}
