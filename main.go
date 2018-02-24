package main

import (
	"github.com/szokodiakos/r8m8/leaderboard"
)

func main() {
	lb := leaderboard.NewService(leaderboard.NewSlackPrinter())
	lb.GetLeaderboard()
}
