package middleware

import (
	"DEEP-backend-image/cerrors"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {
	// header에서 token 추출
	token := c.Get("Authorization")

	if token == "" {
		cerrors.AuthorizationErr("당신은 접근 키를 소지하고 않고 있습니다.")
	} else if token != os.Getenv("ASSIGN_KEY") {
		cerrors.AuthorizationErr("당신은 잘못된 접근 키를 소지하고 있습니다.")
	}

	return c.Next()
}
