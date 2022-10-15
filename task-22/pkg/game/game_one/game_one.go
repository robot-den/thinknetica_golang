// Package game_one contains calculation logic for GameOne
package game_one

import (
	"task-22/pkg/model"
)

// NewGameOne returns a GameOne instance
func NewGameOne() *GameOne {
	game := GameOne{
		Config: model.NewConfig(),
	}
	return &game
}

// GameOne keeps functions to implement game.Calculator interface, and it keeps all state that is required for calculations
type GameOne struct {
	Config *model.Config
}

// Setup reads json config, assings all variables and adds temp storage for iterations
func (g *GameOne) Setup(variant string) {
	model.ReadConfig("game/game_one/GameOne.json", g.Config)
	model.ParseVariant(variant, g.Config)
}

// Iterate iterates selected roundsCount times and fills results
func (g *GameOne) Iterate(roundsCount int) {
	// TODO: implement common iteration logic here (with freespins, respins and any bonus rounds and writings of results
}

// Result returns result of game iterations
func (g *GameOne) Result() *model.CalculationsData {
	return &model.CalculationsData{RoundsCount: 10, TotalWin: 20, TotalBet: 10} // TODO: implement result construction
}
