package middleware

import (
	"time"

	"akubisa/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func AccessLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		latency := time.Since(start).Milliseconds()
		requestID := c.Get("X-Request-ID")

		logger.Logger.Info("access_log",
			"request_id", requestID,
			"method", c.Method(),
			"path", c.Path(),
			"status", c.Response().StatusCode(),
			"latency_ms", latency,
			"client_ip", c.IP(),
		)

		return err
	}
}
