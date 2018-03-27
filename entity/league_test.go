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

	orderedLeaguePlayers := league.GetTop10LeaguePlayers()
	expectedLeaguePlayers := []LeaguePlayer{
		LeaguePlayer{PlayerID: "ProPlayer", Rating: 2200},
		LeaguePlayer{PlayerID: "MediocrePlayer", Rating: 1800},
		LeaguePlayer{PlayerID: "BadPlayer", Rating: 1500},
	}

	assert.Equal(t, orderedLeaguePlayers[0], expectedLeaguePlayers[0])
	assert.Equal(t, orderedLeaguePlayers[1], expectedLeaguePlayers[1])
	assert.Equal(t, orderedLeaguePlayers[2], expectedLeaguePlayers[2])
}
