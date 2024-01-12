package http

import (
	"github.com/gofiber/fiber/v2"

	appError "github.com/dhanielsales/golang-scaffold/internal/error"
)

func ResponseError(c *fiber.Ctx, err error) error {
	if err == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Unknow error",
		})
	}

	appErr, ok := appError.Is(err)
	if ok {
		return c.Status(appErr.Name.Status()).JSON(appErr)
	} else {
		return c.Status(appError.ServerError.Status()).JSON(fiber.Map{
			"name":        appError.ServerError.String(),
			"description": "Unknow error",
		})
	}
}
