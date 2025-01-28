package handlers

import (
	"fiber-boilerplate/app/errors"
	"fiber-boilerplate/app/logger"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type NewsHandler struct {
	Service *service.NewsService
}

func NewNewsHandler(service *service.NewsService) *NewsHandler {
	return &NewsHandler{Service: service}
}

func (h *NewsHandler) CreateNews(c *fiber.Ctx) error {
	var news models.News
	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	if err := h.Service.CreateNews(&news); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create news"})
	}

	return c.JSON(fiber.Map{"message": "News created successfully"})
}

func (h *NewsHandler) GetNewsByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	news, err := h.Service.GetNewsByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "News not found"})
	}

	return c.JSON(news)
}

func (h *NewsHandler) UpdateNews(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var update map[string]interface{}
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid payload"})
	}

	if err := h.Service.UpdateNews(id, update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update news"})
	}

	return c.JSON(fiber.Map{"message": "News updated successfully"})
}

func (h *NewsHandler) DeleteNews(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.Service.DeleteNews(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete news"})
	}

	return c.JSON(fiber.Map{"message": "News deleted successfully"})
}

func (h *NewsHandler) ListNews(c *fiber.Ctx) error {
	newsList, err := h.Service.ListNews()
	if err != nil {
		// Log the error using zap
		logger.Log.Error("Failed to fetch news by ID",
			zap.String("request_id", c.Locals("request_id").(string)),
			zap.Error(err),
			// zap.Stack("stack"),
		)
		return errors.INTERNAL_SERVER_ERROR.Respond(c)
		// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch news"})
	}

	return c.JSON(newsList)
}
