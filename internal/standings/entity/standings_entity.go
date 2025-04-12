package entity

type TeamStanding struct {
	Position      string `json:"position"`
	Team          string `json:"team"`
	Played        string `json:"played"`
	Wins          string `json:"wins"`
	Draws         string `json:"draws"`
	Losses        string `json:"losses"`
	GoalDifference string `json:"goal_difference"`
	Points        string `json:"points"`
}

