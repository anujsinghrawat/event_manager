package middlewares

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/anujsinghrawat/event-manager/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func AuthProctected(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			log.Warnf("Unauthorized request from %s , empty auth header", ctx.IP())

			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  -1,
				"message": "Unauthorized",
				"data":    nil,
			})
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			log.Warnf("Invalid Auth Token from %s", ctx.IP())
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  -2,
				"message": "Unauthorized",
				"data":    nil,
			})
		}
		tokenStr := tokenParts[1]
		secret := []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.GetSigningMethod("HS256").Alg() {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})
		if err != nil || !token.Valid {
			log.Warnf("Invalid Auth Token from %s", ctx.IP())
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  -3,
				"message": "Unauthorized",
				"data":    nil,
			})
		}
		userId := token.Claims.(jwt.MapClaims)["id"]

		if err := db.Model(&models.User{}).Where("id = ?", userId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("User not Found for the given token from %s", ctx.IP())
			return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  -4,
				"message": "Unauthorized",
				"data":    nil,
			})
		}
		ctx.Locals("userId", userId)
		return ctx.Next()
	}
}
