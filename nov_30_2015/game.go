package main

import (
	"log"
)

const (
	Red    = "R"
	Green  = "G"
	Blue   = "U"
	Yellow = "Y"
	Black  = "B"
	White  = "W"

	MaxCodeSize    = 4
	MaxCodeSizeStr = "4"
	MaxGameTurn    = 10

	Failed  = "TIMEOUT"
	Cont    = "CONTINUE"
	Correct = "CORRECT"
)

type Game struct {
	code [MaxCodeSize]string
	turn int
}

func NewGame(keys ...string) *Game {
	// assert
	if len(keys) != MaxCodeSize {
		panic("no of keys should be " + MaxCodeSizeStr)
	}
	for _, color := range keys {
		if color != Red && color != White && color != Black && color != Yellow && color != Green && color != Blue {
			panic("invalid colors")
		}
	}

	return &Game{
		code: [MaxCodeSize]string{keys[0], keys[1], keys[2], keys[3]},
		turn: 0,
	}
}

func (g *Game) Code() [MaxCodeSize]string {
	return g.code
}

func (g *Game) Turn() int {
	return g.turn
}

func (g *Game) Check(guesses ...string) (colorAndPos, colorNoPos int, state string) {
	if len(guesses) != MaxCodeSize {
		panic("no of spot in guess should be " + MaxCodeSizeStr)
	}

	log.Printf("<:%s%s%s%s\\n\n", guesses[0], guesses[1], guesses[2], guesses[3])
	// turn check
	g.turn++
	if g.turn > (MaxGameTurn - 1) {
		return 0, 0, Failed
	}

	colorAndPos, colorNoPos = 0, 0
	unmatchedCode, unmatchedGuesses := make([]string, 0, MaxCodeSize), make([]string, 0, MaxCodeSize)

	// find no of spots matching both color and position first
	// and put the rest to the unmatched list
	for i, color := range g.code {
		if color == guesses[i] {
			colorAndPos++
		} else {
			unmatchedCode = append(unmatchedCode, color)
			unmatchedGuesses = append(unmatchedGuesses, guesses[i])
		}
	}

	// find no of colors from the unmatched list that is present
	// in the guess, but at the wrong position
	for _, color := range unmatchedCode {
		for i, guess := range unmatchedGuesses {
			if color == guess {
				colorNoPos++

				// delete the guess we just processed
				unmatchedGuesses[i] = unmatchedGuesses[len(unmatchedGuesses)-1]
				unmatchedGuesses = unmatchedGuesses[:len(unmatchedGuesses)-1]
				break
			}
		}
	}

	log.Printf("-> colAndPos=%d; colNoPos=%d\n", colorAndPos, colorNoPos)
	if colorAndPos == MaxCodeSize && colorNoPos == 0 {
		return colorAndPos, colorNoPos, Correct
	}

	return colorAndPos, colorNoPos, Cont
}
