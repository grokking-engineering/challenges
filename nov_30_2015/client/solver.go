// this is a solver that randomly pick a guesses
package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	input := bufio.NewReaderSize(os.Stdin, 16)
	var line []byte
	var isPrefix bool
	var err error
	myGame := NewGame()

	for {
		// assert
		line, isPrefix, err = input.ReadLine()
		log.Println(string(line))
		switch {
		case isPrefix:
			panic("driver sends a big message")
		case err != nil: // no need to handle EOF, as driver will guarantee that
			panic("cannot read from stdin; err= " + err.Error())
		}

		// some pre-processing to take care of the START
		switch {
		case string(line) == "START":
			myGame.Restart()
			myGame.Process(0, 0, 0, "CONTINUE")
		default:
			turn, colorAndPos, colorNoPos, state := parseInput(line)
			myGame.Process(turn, colorAndPos, colorNoPos, state)
		}
	}
}

func parseInput(line []byte) (turn, colorAndPos, colorNoPos int, state string) {
	var err error
	elems := strings.Split(string(line), " ")
	if len(elems) != 4 {
		panic("bad driver input")
	}

	turn, err = strconv.Atoi(elems[0])
	if err != nil {
		panic("turn is not a number")
	}
	colorAndPos, err = strconv.Atoi(elems[1])
	if err != nil {
		panic("colorAndPos is not a number")
	}
	colorNoPos, err = strconv.Atoi(elems[2])
	if err != nil {
		panic("colorNoPos is not a number")
	}
	state = elems[3]
	if state != "CONTINUE" && state != "TIMEOUT" {
		panic("unknown game state")
	}
	return turn, colorAndPos, colorNoPos, elems[3]
}
