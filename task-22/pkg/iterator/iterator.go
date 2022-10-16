package iterator

import "log"

type Game interface {
	// SetupStateBeforeRound is a hook where you can setup (or update) all data you need later
	SetupStateBeforeRound()
	// GenerateScreenInSpin is a hook where you can generate screen in way you need (or use skrn package)
	GenerateScreenInSpin()
	// CalculateWinInSpin is a hook where you can calculate win from spin screen (or use calcwin package)
	CalculateWinInSpin()
	// SetupStateAfterSpin is a hook where you can clean (or update) all data you need before Freespin or Bonus game started
	SetupStateAfterSpin()
	// SetupStateAfterRound is a hook where you can clean (or update) all data you need (it's called after full round)
	SetupStateAfterRound()
}

type GameWithFreespins interface {
	SetupStateBeforeFreespinPack()
	FreespinsGiven() int
	SetupStateBeforeFreespin()
	GenerateScreenInFreespin()
	CalculateWinInFreespin()
	SetupStateAfterFreespin()
	SetupStateAfterFreespinPack()
}

// GameWithBonus suggests that game will implement full bonus game cycle by itself
type GameWithBonus interface {
	ProcessBonusGame()
}

func Iterate(gInterface interface{}, roundsCount int) {
	g, ok := gInterface.(Game)
	if !ok {
		log.Fatal("can't convert passed game into iterator.game")
	}

	roundsDone := 0

	for roundsCount > roundsDone {
		g.SetupStateBeforeRound()
		g.GenerateScreenInSpin()
		g.CalculateWinInSpin()
		g.SetupStateAfterSpin()

		if gWithFs, ok := g.(GameWithFreespins); ok {
			ProcessFreespinGame(gWithFs)
		}

		if gWithBonus, ok := g.(GameWithBonus); ok {
			gWithBonus.ProcessBonusGame()
		}

		g.SetupStateAfterRound()

		roundsDone++
	}
}

func ProcessFreespinGame(g GameWithFreespins) {
	g.SetupStateBeforeFreespinPack()
	freespinsDone := 0
	for g.FreespinsGiven() > freespinsDone {
		g.SetupStateBeforeFreespin()
		g.GenerateScreenInFreespin()
		g.CalculateWinInFreespin()
		g.SetupStateAfterFreespin()

		freespinsDone++
	}

	g.SetupStateAfterFreespinPack()
}
