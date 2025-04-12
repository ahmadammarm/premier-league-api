package scrapper

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/ahmadammarm/premier-league-api/internal/standings/entity"
)

type PremierLeagueScrapper struct {
	url string
}

func NewPremierLeagueScrapper(url string) *PremierLeagueScrapper {
	return &PremierLeagueScrapper{
		url: url,
	}
}

func (scrapper *PremierLeagueScrapper) FetchStandings() ([]entity.TeamStanding, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Get(scrapper.url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch standings: %s", response.Status)
	}


	document, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		return nil, err
	}

	var standings []entity.TeamStanding

	document.Find("li.Standing_standings__row__5sdZG").Each(func(i int, row *goquery.Selection) {
		positionElem := row.Find("div.Standing_standings__cell__5Kd0W").First()
		teamElem := row.Find("p.Standing_standings__teamName__psv61")
		stats := row.Find("div.Standing_standings__cell__5Kd0W")

		if positionElem.Length() > 0 && teamElem.Length() > 0 && stats.Length() >= 7 {
			team := entity.TeamStanding{
				Position:       strings.TrimSpace(positionElem.Text()),
				Team:           strings.TrimSpace(teamElem.Text()),
				Played:         strings.TrimSpace(stats.Eq(2).Text()),
				Wins:           strings.TrimSpace(stats.Eq(3).Text()),
				Draws:          strings.TrimSpace(stats.Eq(4).Text()),
				Losses:         strings.TrimSpace(stats.Eq(5).Text()),
				GoalDifference: strings.TrimSpace(stats.Eq(6).Text()),
				Points:         strings.TrimSpace(stats.Eq(7).Text()),
			}
			standings = append(standings, team)
		}
	})

	if len(standings) == 0 {
		return nil, fmt.Errorf("no standings data found")
	}

	return standings, nil

}
