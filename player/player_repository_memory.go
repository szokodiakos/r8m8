package player

import (
	"github.com/szokodiakos/r8m8/entity"
	"github.com/szokodiakos/r8m8/transaction"
)

// PlayerRepositoryMemory struct
type PlayerRepositoryMemory struct {
	Players []entity.Player
}

// Add func
func (p *PlayerRepositoryMemory) Add(tr transaction.Transaction, player entity.Player) (entity.Player, error) {
	p.Players = append(p.Players, player)
	return player, nil
}

// GetMultipleByIDs func
func (p *PlayerRepositoryMemory) GetMultipleByIDs(tr transaction.Transaction, ids []string) ([]entity.Player, error) {
	players := []entity.Player{}
	for i := range p.Players {
		for j := range ids {
			if p.Players[i].ID == ids[j] {
				players = append(players, p.Players[i])
			}
		}
	}
	return players, nil
}

// GetByID func
func (p *PlayerRepositoryMemory) GetByID(tr transaction.Transaction, id string) (entity.Player, error) {
	player := entity.Player{}
	for i := range p.Players {
		if p.Players[i].ID == id {
			player = p.Players[i]
		}
	}
	return player, nil
}
