// Package game_one. The spin file contains all hooks that are required to implement iterator.Game
package game_one

import "fmt"

func (g *GameOne) SetupStateBeforeRound() {
	fmt.Println("R")
}

func (g *GameOne) GenerateScreenInSpin() {
	fmt.Println("I")
}

func (g *GameOne) CalculateWinInSpin() {
	fmt.Println("T")
}

func (g *GameOne) SetupStateAfterSpin() {
	fmt.Println("A")
}

func (g *GameOne) SetupStateAfterRound() {
	fmt.Println("<3")
}
