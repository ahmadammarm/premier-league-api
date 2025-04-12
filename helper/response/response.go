package response

import (
	"github.com/ahmadammarm/premier-league-api/standings"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message,omitempty"`
	Data    []standings.TeamStanding `json:"data,omitempty"`
}

func SendResponse(context *fiber.Ctx, message string, data []standings.TeamStanding) error {
	return context.JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}
