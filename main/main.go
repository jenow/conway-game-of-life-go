package main

import "time"

const (
	WIDTH  = 50
	HEIGHT = 50
	SEED   = 8
	CYCLES = 3000
)

func main() {
	g := NewGame()

	for i := 0; i < CYCLES; i++ {
		print("\033[H\033[2J")
		g.Display()
		g.Live()
		time.Sleep(time.Second / 30)
	}
}
