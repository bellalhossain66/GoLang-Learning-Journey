package handler

import(
	"github.com/gofiber/fiber/v2"
	"go-fiber-mvc-app/service"
	"go-fiber-mvc-app/dto"
)

func LoginUser(c *fiber.Ctx) error {
	var body struct {
		Username string `json:username`
		Password string `json:password`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(dto.ErrorWithMessage[string](nil, "Invalid request"))
	}

	token, err := service.LoginUser(body.Username, body.Password);
	if err != nil {
		return c.Status(401).JSON(dto.ErrorWithMessage[string](nil, err.Error()))
	}

	return c.JSON(dto.OKWithMessage[string](&token, "Login successful"))
}