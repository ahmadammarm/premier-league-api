package handler

import (
	"strings"

	"github.com/ahmadammarm/premier-league-api/helper/response"
	"github.com/ahmadammarm/premier-league-api/internal/standings/service"
	"github.com/gofiber/fiber/v2"
)

type StandingsHandler struct {
	service service.StandingService
}

func NewStandingsHandler(service service.StandingService) *StandingsHandler {
	return &StandingsHandler{
		service: service,
	}
}

func (handler *StandingsHandler) GetAllStandings(context *fiber.Ctx) error {
	standings, err := handler.service.GetAllStandings()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch standings: "+err.Error())
	}

	response := response.NewStandingsResponse(true, "Successfully fetched standings", standings)
	return context.JSON(response)
}

func (handler *StandingsHandler) GetTeamByPosition(context *fiber.Ctx) error {
	position := context.Params("position")

	standings, err := handler.service.GetTeamsByPosition(position)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := response.NewStandingsResponse(true, "Successfully fetched team at position "+position, standings)
	return context.JSON(response)
}

func (handler *StandingsHandler) GetTeamByName(context *fiber.Ctx) error {
	name := context.Params("name")

	standings, err := handler.service.GetTeamByName(name)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := response.NewStandingsResponse(
		true,
		"Successfully found teams matching '"+name+"'",
		standings,
	)
	return context.JSON(response)
}

func (handler *StandingsHandler) GetTeamsByZone(context *fiber.Ctx) error {
	zone := context.Params("zone")

	standings, err := handler.service.GetTeamsByZone(zone)
	if err != nil {
		if strings.Contains(err.Error(), "invalid zone") {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	response := response.NewStandingsResponse(
		true,
		"Successfully fetched teams in "+zone+" zone",
		standings,
	)
	return context.JSON(response)
}
