// terminology
// ----------------------------------------------------
// 	- We have the this main program (driver), and your subprogram (solver)
//	- They talk via pipes:
//    - driver's stdout > solver's stdin
//    - solver's stdout > driver's stdin
//  - All commands are terminated with the character '\n' (your separator)
//  - To make it simpler to write examples, whenever you see:
//		- '>:' it indicates that driver is writing into solver's stdin
//		- '<:' it indicates that solver is writing to driver's stdin
//		- The data written is on the right of the ':'
//	- When we test your program, we expect it to terminates when the game end
//		e.g, when you receive either TIMEOUT or CORRECT
//	- After you have writter the guesses to stdout, flush your output buffer
//		so we can read it
//
// convention
// ----------------------------------------------------
// colors are indicated using 1 ASCII character, as followed:
// 	- Red: 		R
// 	- Green: 	G
// 	- Blue: 	U
// 	- Yellow: Y
// 	- Black: 	B
// 	- White: 	W
//
// input/output
// ----------------------------------------------------
// 	- >:START\n						// indicate new game
//  - <:RRRR\n						// solver make a first guess
//	- >:0 1 2 CONTINUE\n	// driver returns the following 4 things:
//												// 		<turn no>
//												//		<no of spot matching color & position (1)>
//												//		<no of spot matching color ONLY, excluding (1)>
//												// 		<a state>, where
//												//			CONTINUE indicate you can make another guess
//												// 			CORRECT indicate you have guessed correctly
//												//			TIMEOUT indicate you didn't find the answer
//
// example: let's say the code is RWBY
// ----------------------------------------------------
// >:START\n
// <:RRRR\n
// >:1 1 0 CONTINUE\n			// 1st spot matched
// <:RBBB\n
// >:2 2 0 CONTINUE\n     // 1st & 3rd spot matched
// <:RBWW\n
// >:3 1 2 CONTINUE\n			// 1st spot matched, 2nd & 3rd are in wrong position
// <:RWBB\n
// >:4 3 0 CONTINUE\n     // 1st, 2nd & 3rd matched
// <:RWBY\n
// >:5 4 0 CORRECT\n      // all matched
//
//
// the same game might exceed 10 turns, in which case we return a TIMEOUT
//
// ... some other steps ..
// <:RWBB\n
// >:9 3 0 CONTINUE\n			// 1st, 2nd & 3rd matched
// <:RWBU\n
// >:10 3 0 TIMEOUT				// 1st, 2nd & 3rd matched, timeout
//
package main

import (
	"flag"
	"log"
	"sync"
	"sync/atomic"
)

var (
	codeFile string
	showHelp bool
)

func init() {
	flag.StringVar(&codeFile, "c", "", "path to the file containing codes for the game")
	flag.BoolVar(&showHelp, "h", false, "show usage")
}

func main() {
	flag.Parse()

	if showHelp {
		printHelp()
		return
	}

	assertParams()

	// init
	// ---------------------------------------------------------
	allCodes := readAllCodes(codeFile)
	var wg sync.WaitGroup
	var correct uint32
	correct = 0

	// run all tests
	// ---------------------------------------------------------
	for _, code := range allCodes {
		wg.Add(1)

		// copy to avoid race conditions
		var gameCode [MaxCodeSize]string
		for i := 0; i < MaxCodeSize; i++ {
			gameCode[i] = code[i]
		}

		// run solver in parallel
		go func() {
			matched := runTest(gameCode)
			defer wg.Done()
			if matched {
				atomic.AddUint32(&correct, 1)
			}
		}()
	}
	wg.Wait()
	log.Printf("Done! Result: %d correct / %d total\n", correct, len(allCodes))
}

// run a game with the solver, return true if it figures out the code
func runTest(code [MaxCodeSize]string) bool {
	log.Printf("new game: %s%s%s%s\n", code[0], code[1], code[2], code[3])

	// setup
	// ---------------------------------------------------------
	game := NewGame(code[0], code[1], code[2], code[3])
	solver := getSolver()
	solverStdin, err := solver.StdinPipe()
	if err != nil {
		panic("cannot get solver's stdin")
	}
	solverStdout, err := solver.StdoutPipe()
	if err != nil {
		panic("cannot get solver's stdout")
	}
	input := getClientInput(solverStdout)

	// start game
	// ---------------------------------------------------------
	solver.Start()

	// tell it to start
	log.Printf(">:START\\n\n")
	solverStdin.Write([]byte("START\n"))

	// keep taking solver's output and give its new result
	for item := range input {
		log.Printf("<:%s%s%s%s\\n\n", item[0], item[1], item[2], item[3])

		colorAndPos, colorNoPos, state := game.Check(item[0], item[1], item[2], item[3])
		writeToSolver(solverStdin, game.Turn(), colorAndPos, colorNoPos, state)

		switch {
		case state == Failed:
			return false
		case state == Correct:
			return true
		}
	}

	// should not come here
	return false
}
