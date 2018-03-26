package entity

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetTop10LeaguePlayers(t *testing.T) {
	league := League{
		LeaguePlayers: []LeaguePlayer{
			LeaguePlayer{
				PlayerID: "MediocrePlayer",
				Rating:   1800,
			},
			LeaguePlayer{
				PlayerID: "BadPlayer",
				Rating:   1500,
			},
			LeaguePlayer{
				PlayerID: "ProPlayer",
				Rating:   2200,
			},
		},
	}

	orderedLeaguePlayers := league.GetTop10LeaguePlayers()
	if isLeaguePlayersNotInOrder(orderedLeaguePlayers) {
		t.Error("They should be in order", spew.Sdump(orderedLeaguePlayers))
	}
}

func isLeaguePlayersNotInOrder(orderedLeaguePlayers []LeaguePlayer) bool {
	isFirstInPlace := orderedLeaguePlayers[0].PlayerID == "ProPlayer"
	isSecondInPlace := orderedLeaguePlayers[1].PlayerID == "MediocrePlayer"
	isThirdInPlace := orderedLeaguePlayers[2].PlayerID == "BadPlayer"

	return !(isFirstInPlace && isSecondInPlace && isThirdInPlace)
}
