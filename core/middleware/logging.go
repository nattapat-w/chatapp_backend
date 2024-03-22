package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Logging struct {
}

func LoggingMiddleware(c *fiber.Ctx) error {
	// Start timer
	start := time.Now()

	// Process request
	err := c.Next()

	// Calculate processing time
	duration := time.Since(start)

	// Log the information
	fmt.Printf("Request URL: %s - Method: %s - Duration: %s\n", c.OriginalURL(), c.Method(), duration)

	return err
}
