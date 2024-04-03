package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/nattapat-w/chatapp/core/user/dto"
)

func CheckToken(c *fiber.Ctx) error {
	// Get token from cookie
	var cookie string
	cookie = c.Cookies("jwt")
	cookie2 := c.Query("token")
	if cookie2 != "" {
		cookie = cookie2
	}
	fmt.Println(cookie)
	if cookie == "" {
		// Token not found
		return c.Status(http.StatusUnauthorized).SendString("Unauthorized token not found")
	}

	// Parse JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil && token.Valid {
		return c.Status(http.StatusUnauthorized).SendString("Unauthorized token is invalid")
	}
	payload := token.Claims.(jwt.MapClaims)
	userData := &dto.UserDataDTO{
		Username:    payload["username"].(string),
		ID:          uint(payload["user_id"].(float64)),
		DisplayName: payload["displayName"].(string),
	}
	c.Locals("userData", userData)

	return c.Next()
}
