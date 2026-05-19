package response

import "github.com/gofiber/fiber/v2"

type envelope struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func OK(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(envelope{Success: true, Data: data})
}

func Created(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(envelope{Success: true, Data: data})
}

func Message(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusOK).JSON(envelope{Success: true, Message: msg})
}

func BadRequest(c *fiber.Ctx, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(envelope{Success: false, Error: err})
}

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(envelope{Success: false, Error: "unauthorized"})
}

func Forbidden(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(envelope{Success: false, Error: "forbidden"})
}

func NotFound(c *fiber.Ctx, resource string) error {
	return c.Status(fiber.StatusNotFound).JSON(envelope{Success: false, Error: resource + " not found"})
}

func Conflict(c *fiber.Ctx, err string) error {
	return c.Status(fiber.StatusConflict).JSON(envelope{Success: false, Error: err})
}

func InternalError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(envelope{Success: false, Error: "internal server error"})
}
