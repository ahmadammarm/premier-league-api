package main

import (
	"log"

	"github.com/ahmadammarm/premier-league-api/config"
	"github.com/ahmadammarm/premier-league-api/helper/errorhandler"
	"github.com/ahmadammarm/premier-league-api/helper/scrapper"
	"github.com/ahmadammarm/premier-league-api/internal/standings/handler"
	"github.com/ahmadammarm/premier-league-api/internal/standings/routes"
	"github.com/ahmadammarm/premier-league-api/internal/standings/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	congiguration := config.Configuration()

	premierLeagueScrapper := scrapper.NewPremierLeagueScrapper(congiguration.ScrapperURL)

	standingsService := service.NewStandingsService(premierLeagueScrapper)

	standingsHandler := handler.NewStandingsHandler(standingsService)

	application := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          errorhandler.CustomErrorHandler,
	})

	routes.RegisterRoutes(application, standingsHandler)

	log.Printf("Starting Premier League Standings API on port %s...", congiguration.Port)
	log.Fatal(application.Listen(":" + congiguration.Port))

}


