// Package game_selector dispatch calculation logic based on selected game
package game_selector

import (
	"log"
	"task-22/pkg/game"
	"task-22/pkg/game/game_one"
	"task-22/pkg/model"
)

func CalculateGame(gameId, variant string, roundsCount int) *model.CalculationsData {
	var calculator game.Calculator
	switch gameId {
	case "GameOne":
		calculator = game_one.NewGameOne()
	default:
		log.Fatal("provided game is not implemented")
	}

	calculator.Setup(variant)
	calculator.Iterate(roundsCount)
	return calculator.Result()
}
