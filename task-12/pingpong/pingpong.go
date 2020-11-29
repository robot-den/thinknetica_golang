package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type game struct {
	scoreboard map[string]int
	limit      int
	ch         chan string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := newGame(10)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go play("Player #1", game, wg)
	go play("Player #2", game, wg)

	game.ch <- "begin"
	wg.Wait()
	game.printScores()
}

func play(name string, game *game, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		hit := "ping"
		competitorHit, ok := <-game.ch
		if competitorHit == "ping" {
			hit = "pong"
		}

		if game.isDone() {
			if ok {
				close(game.ch)
			}
			break
		}

		if game.randWin() {
			game.incrScore(name)
			hit = "stop"
		}

		fmt.Printf("%s: %s\n", name, hit)
		game.ch <- hit
	}
}

func newGame(limit int) *game {
	s := game{
		limit:      limit,
		scoreboard: make(map[string]int),
		ch:         make(chan string),
	}
	return &s
}

func (g *game) incrScore(playerName string) {
	g.scoreboard[playerName] += 1
}

func (g *game) isDone() bool {
	for _, v := range g.scoreboard {
		if v >= g.limit {
			return true
		}
	}
	return false
}

func (g *game) randWin() bool {
	n := rand.Intn(100)
	if n < 20 {
		return true
	}
	return false
}

func (g *game) printScores() {
	fmt.Println("\nGame just end!")
	for player, score := range g.scoreboard {
		fmt.Printf("%s: %d\n", player, score)
	}
}
