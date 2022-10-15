// Package main only reads params from console, runs calculations through game selector and then saves result via storage package
package main

import (
	"flag"
	"log"
	"task-22/pkg/game_selector"
)

// flags that calculator accepts from console
var flagGameId string
var flagGameVariant string
var flagRoundsCount int
var flagStorageType string

// main reads flags, runs calculations via GameSelector and saves results via Storage
func main() {
	parseFlags()
	result := game_selector.CalculateGame(flagGameId, flagGameVariant, flagRoundsCount)
	log.Println(result) // TODO: replace it with pretty print and savings logic
}

// parseFlags parses flags from the console and stores them in variables
func parseFlags() {
	flag.StringVar(&flagGameId, "game", "", "the game identifier without variant postfixes")
	flag.StringVar(&flagGameVariant, "variant", "", "format '<mode>:<math>', example: base:97")
	flag.IntVar(&flagRoundsCount, "rounds", 0, "rounds count to process")
	flag.StringVar(&flagStorageType, "storage", "json", "where to store results, available 'json' and 'print'")
	flag.Parse()

	if len(flagGameVariant) == 0 {
		log.Fatal("you must provide game variant, pass --help for more")
	}

	if len(flagGameId) == 0 {
		log.Fatal("you must provide game identifier, pass --help for more")
	}

	if flagRoundsCount == 0 {
		log.Fatal("you must provide rounds count, pass --help for more")
	}
}
