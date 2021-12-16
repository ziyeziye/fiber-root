package middleware

import (
	"fiber-root/app/api"
	"fiber-root/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func Jwt() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Get("token")
		if token == "" {
			c.Status(fiber.StatusBadRequest)
			return api.Response(c, fiber.ErrBadRequest, nil)
		} else {
			claims, err := jwt.ParseToken(token)
			if err != nil {
				c.Status(fiber.StatusUnauthorized)
				return api.Response(c, fiber.ErrUnauthorized, nil)
			} else if time.Now().Unix() > claims.ExpiresAt {
				c.Status(fiber.StatusExpectationFailed)
				return api.Response(c, fiber.ErrExpectationFailed, nil)
			}
		}
		return c.Next()
	}
}
