package routes

import (
	"github.com/ahmadammarm/premier-league-api/internal/standings/handler"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, sh *handler.StandingsHandler) {
	api := app.Group("/api")

	standings := api.Group("/standings")
	standings.Get("/", sh.GetAllStandings)
	standings.Get("/position/:position", sh.GetTeamByPosition)
	standings.Get("/team/:name", sh.GetTeamByName)
	standings.Get("/zone/:zone", sh.GetTeamsByZone)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}