package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTop10LeaguePlayers(t *testing.T) {
	league := League{
		LeaguePlayers: []LeaguePlayer{
			LeaguePlayer{PlayerID: "MediocrePlayer", Rating: 1800},
			LeaguePlayer{PlayerID: "BadPlayer", Rating: 1500},
			LeaguePlayer{PlayerID: "ProPlayer", Rating: 2200},
		},
	}
	expectedLeaguePlayers := []LeaguePlayer{
		LeaguePlayer{PlayerID: "ProPlayer", Rating: 2200},
		LeaguePlayer{PlayerID: "MediocrePlayer", Rating: 1800},
		LeaguePlayer{PlayerID: "BadPlayer", Rating: 1500},
	}
	testGetTop10LeaguePlayers(t, expectedLeaguePlayers, league)
}

func TestGetTop10LeaguePlayersWhenRatingsAreEqual(t *testing.T) {
	league := League{
		LeaguePlayers: []LeaguePlayer{
			LeaguePlayer{PlayerID: "MediocrePlayer", Rating: 1500, winCount: 5},
			LeaguePlayer{PlayerID: "BadPlayer", Rating: 1500, winCount: 0},
			LeaguePlayer{PlayerID: "ProPlayer", Rating: 1500, winCount: 10},
		},
	}
	expectedLeaguePlayers := []LeaguePlayer{
		LeaguePlayer{PlayerID: "ProPlayer", Rating: 1500, winCount: 10},
		LeaguePlayer{PlayerID: "MediocrePlayer", Rating: 1500, winCount: 5},
		LeaguePlayer{PlayerID: "BadPlayer", Rating: 1500, winCount: 0},
	}
	testGetTop10LeaguePlayers(t, expectedLeaguePlayers, league)
}

func TestGetTop10LeaguePlayersWhenRatingsAndWinCountsAreEqual(t *testing.T) {
	league := League{
		LeaguePlayers: []LeaguePlayer{
			LeaguePlayer{PlayerID: "MediocrePlayer", Rating: 1500, winCount: 10, matchCount: 20},
			LeaguePlayer{PlayerID: "BadPlayer", Rating: 1500, winCount: 10, matchCount: 30},
			LeaguePlayer{PlayerID: "ProPlayer", Rating: 1500, winCount: 10, matchCount: 10},
		},
	}
	expectedLeaguePlayers := []LeaguePlayer{
		LeaguePlayer{PlayerID: "ProPlayer", Rating: 1500, winCount: 10, matchCount: 10},
		LeaguePlayer{PlayerID: "MediocrePlayer", Rating: 1500, winCount: 10, matchCount: 20},
		LeaguePlayer{PlayerID: "BadPlayer", Rating: 1500, winCount: 10, matchCount: 30},
	}
	testGetTop10LeaguePlayers(t, expectedLeaguePlayers, league)
}

func testGetTop10LeaguePlayers(t *testing.T, expectedLeaguePlayers []LeaguePlayer, league League) {
	orderedLeaguePlayers := league.GetTop10LeaguePlayers()
	assertLeaguePlayersInOrder(t, expectedLeaguePlayers, orderedLeaguePlayers)
}

func assertLeaguePlayersInOrder(t *testing.T, expectedLeaguePlayers []LeaguePlayer, actualLeaguePlayers []LeaguePlayer) {
	assert.Equal(t, len(expectedLeaguePlayers), len(actualLeaguePlayers))
	for i := range expectedLeaguePlayers {
		assert.Equal(t, expectedLeaguePlayers[i], actualLeaguePlayers[i])
	}
}
