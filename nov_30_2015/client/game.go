package main

import (
	"fmt"
	"log"
	"math/rand"
)

var (
	choices = []string{"R", "G", "U", "Y", "B", "W"}
)

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Restart() {
	log.Println("starting new game")
}

func (g *Game) Process(turn, colorAndPos, colorNoPos int, state string) {
	log.Printf("[RECEIVED]: turn= %d; colorAndPos= %d; colorNoPos= %d; state=%s\n", turn, colorAndPos, colorNoPos, state)

	if state == "TIMEOUT" {
		log.Printf("I have failed\n")
	} else if state == "CORRECT" {
		log.Printf("I have succeeded\n")
	} else {
		writeToDriver()
	}
}

func writeToDriver() {
	guess := choices[rand.Intn(4)] + choices[rand.Intn(4)] + choices[rand.Intn(4)] + choices[rand.Intn(4)]
	// guess := "RWBY"
	log.Printf("[ SENDING]: %s\n", guess)
	fmt.Printf("%s\n", guess)
}
