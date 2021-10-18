package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/oinpentuls/auth-with-fiber/repositories"
)

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	log.Println(cookie)
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	getUser, err := repositories.GetUserById(claims.Issuer)

	if err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(getUser)

}
