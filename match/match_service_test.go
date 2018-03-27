package match

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/rating"
)

func TestCalculatePlayerChanges(t *testing.T) {
	matchService := NewService(rating.NewStrategyElo())

	repoLeaguePlayers := []entity.LeaguePlayer{
		entity.LeaguePlayer{PlayerID: "Winner Player", Rating: 1500},
		entity.LeaguePlayer{PlayerID: "Player who did not participate", Rating: 2000},
		entity.LeaguePlayer{PlayerID: "Loser Player", Rating: 1500},
	}

	players := []entity.Player{
		entity.Player{ID: "Winner Player", DisplayName: "irrelevant"},
		entity.Player{ID: "Loser Player", DisplayName: "irrelevant2"},
	}

	adjustedLeaguePlayers, createdMatchPlayers := matchService.CalculatePlayerChanges(repoLeaguePlayers, players)

	expectedLeaguePlayers := []entity.LeaguePlayer{
		entity.LeaguePlayer{PlayerID: "Winner Player", Rating: 1516},
		entity.LeaguePlayer{PlayerID: "Player who did not participate", Rating: 2000},
		entity.LeaguePlayer{PlayerID: "Loser Player", Rating: 1484},
	}

	expectedMatchPlayers := []entity.MatchPlayer{
		entity.MatchPlayer{PlayerID: "Winner Player", HasWon: true, RatingChange: 16},
		entity.MatchPlayer{PlayerID: "Loser Player", HasWon: false, RatingChange: -16},
	}

	assert.ElementsMatch(t, expectedLeaguePlayers, adjustedLeaguePlayers)
	assert.ElementsMatch(t, expectedMatchPlayers, createdMatchPlayers)
}
