package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Map    [WIDTH][HEIGHT]bool
	TmpMap [WIDTH][HEIGHT]bool
}

func NewGame() *Game {
	g := &Game{}
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < WIDTH*HEIGHT/SEED; i++ {
		g.Map[rand.Intn(WIDTH)][rand.Intn(HEIGHT)] = true
	}

	g.TmpMap = g.Map

	return g
}

func (g Game) Display() {
	for j := 0; j < HEIGHT; j++ {
		for i := 0; i < WIDTH; i++ {
			if g.Map[j][i] == true {
				fmt.Printf("o")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func (g Game) isAlive(i, j int) bool {
	i += WIDTH
	i %= WIDTH
	j += HEIGHT
	j %= HEIGHT
	return g.Map[i][j]
}

func (g Game) getNextState(i, j int) bool {
	alive := 0
	for k := -1; k <= 1; k++ {
		for l := -1; l <= 1; l++ {
			if (j != 0 || i != 0) && g.isAlive(i+k, j+l) {
				alive++
			}
		}
	}

	return alive == 3 || alive == 2 && g.isAlive(i, j)
}

func (g *Game) Live() {
	for j := 0; j < HEIGHT; j++ {
		for i := 0; i < WIDTH; i++ {
			g.TmpMap[i][j] = g.getNextState(i, j)
		}
	}

	g.Map = g.TmpMap
}
