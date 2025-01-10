package middleware

import (
	"encoding/json"
	"fiber-boilerplate/app/logger"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func LogMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		log := logger.NewLogger()

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		body := c.Body()
		var bodyMap map[string]interface{}

		json.Unmarshal(body, &bodyMap)

		log.Logger.WithFields(logrus.Fields{
			"timestamp": timestamp,
			"method":    c.Method(),
			"path":      c.Path(),
			// "body":      fiber.Map(bodyMap),
			// "body": string(body),
		}).Info("Request received")
		if body != nil {
			// fmt.Printf("\x1b[34m[%s] Request received\x1b[0m\n", timestamp)
			// fmt.Printf("\x1b[34mMethod:\x1b[0m %s\n", c.Method())
			// fmt.Printf("\x1b[34mPath:\x1b[0m %s\n", c.Path())
			fmt.Printf("\x1b[34mBody:\x1b[0m ")
			bodyJSON, _ := json.MarshalIndent(bodyMap, "", "  ")
			fmt.Printf("%s\n", bodyJSON)
		}
		return c.Next()
	}
}
