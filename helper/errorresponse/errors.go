package errorresponse

import (
	"log"

	"github.com/ahmadammarm/premier-league-api/helper/response"
	"github.com/gofiber/fiber/v2"
)

func SendErrorResponse(context *fiber.Ctx, message string, err error) error {
	log.Printf("Error: %s - %v", message, err)
	context.Status(500)
	return context.JSON(response.Response{
		Success: false,
		Message: message + ": " + err.Error(),
	})
}
