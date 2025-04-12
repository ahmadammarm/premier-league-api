package response

import (
	"time"

	"github.com/ahmadammarm/premier-league-api/internal/standings/entity"
)

type StandingsResponse struct {
	Success   bool           `json:"success"`
	Message   string         `json:"message,omitempty"`
	Timestamp string         `json:"timestamp"`
	Standings []entity.TeamStanding `json:"standings,omitempty"`
}

func NewStandingsResponse(success bool, message string, standings []entity.TeamStanding) StandingsResponse {
	return StandingsResponse{
		Success:   success,
		Message:   message,
		Timestamp: time.Now().Format(time.RFC3339),
		Standings: standings,
	}
}