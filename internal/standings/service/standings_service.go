package service

import (
	"fmt"
	"strings"

	"github.com/ahmadammarm/premier-league-api/internal/standings/entity"
)

type StandingScrapper interface {
	FetchStandings() ([]entity.TeamStanding, error)
}

type StandingService interface {
	GetAllStandings() ([]entity.TeamStanding, error)
	GetTeamByPosition(position string) ([]entity.TeamStanding, error)
	GetTeamByName(name string) ([]entity.TeamStanding, error)
	GetTeamByZone(zone string) ([]entity.TeamStanding, error)
}

type standingService struct {
	scrapper StandingScrapper
}

// GetAllStandings
func (service *standingService) GetAllStandings() ([]entity.TeamStanding, error) {
	standings, err := service.scrapper.FetchStandings()
	if err != nil {
		return nil, err
	}
	return standings, nil
}

// GetTeamByPosition
func (service *standingService) GetTeamByPosition(position string) ([]entity.TeamStanding, error) {
	standings, err := service.GetAllStandings()
	if err != nil {
		return nil, err
	}

	for _, team := range standings {
		if team.Position == position {
			return []entity.TeamStanding{team}, nil
		}
	}

	return nil, fmt.Errorf("no team found at position %s", position)
}

// GetTeamByName
func (service *standingService) GetTeamByName(name string) ([]entity.TeamStanding, error) {
	standings, err := service.GetAllStandings()
	if err != nil {
		return nil, err
	}

	name = strings.ToLower(name)
	var result []entity.TeamStanding

	for _, team := range standings {
		if strings.Contains(strings.ToLower(team.Team), name) {
			result = append(result, team)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no team found with name containing '%s'", name)
	}

	return result, nil
}

// GetTeamByZone
func (service *standingService) GetTeamByZone(zone string) ([]entity.TeamStanding, error) {
	standings, err := service.GetAllStandings()
	if err != nil {
		return nil, err
	}

	zone = strings.ToLower(zone)
	var result []entity.TeamStanding

	switch zone {
	case "champions":
		if len(standings) >= 5 {
			result = standings[:5]
		} else {
			result = standings
		}
	case "europa":
		if len(standings) >= 7 {
			result = standings[5:7]
		} else if len(standings) > 5 {
			result = standings[5:]
		}
	case "conference":
		if len(standings) >= 8 {
			result = standings[7:8]
		}
	case "relegation":
		if len(standings) >= 3 {
			result = standings[len(standings)-3:]
		}
	default:
		return nil, fmt.Errorf("invalid zone: %s. Available zones: champions, europa, conference, relegation", zone)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no teams found in zone '%s'", zone)
	}

	return result, nil
}

func NewStandingsService(scrapper StandingScrapper) StandingService {
	return &standingService{
		scrapper: scrapper,
	}
}
