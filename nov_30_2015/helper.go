package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func assertParams() {
	if len(codeFile) < 1 {
		panic("no code file")
	}

	if len(flag.Args()) < 1 {
		panic("no solver")
	}
}

func printHelp() {
	fmt.Println(`codebreaker [-h] -c=file <commands>
  -h               Show usage
  -c               Path to file containing codes
  <commands>       Command to run your solver`)
}

func readAllCodes(name string) [][MaxCodeSize]string {
	file, err := os.Open(name)
	if err != nil {
		panic("cannot open code file")
	}
	allCodes, err := ioutil.ReadAll(file)
	if err != nil {
		panic("cannot read code file content")
	}

	result := make([][MaxCodeSize]string, 0, 50)
	codes := strings.Split(string(allCodes), "\n")
	for _, code := range codes {
		result = append(result, [MaxCodeSize]string{toVal(code[0]), toVal(code[1]), toVal(code[2]), toVal(code[3])})
	}

	return result
}

func toVal(color byte) string {
	switch color {
	case 'R':
		return Red
	case 'G':
		return Green
	case 'U':
		return Blue
	case 'Y':
		return Yellow
	case 'B':
		return Black
	case 'W':
		return White
	}

	panic("unknown color")
}

// keep reading the solver's output and put it into a channel
func getClientInput(inputFile io.Reader) chan [MaxCodeSize]string {
	result := make(chan [MaxCodeSize]string)

	go func() {
		input := bufio.NewReaderSize(inputFile, 16)
		var line []byte
		var isPrefix bool
		var err error
		var guesses [MaxCodeSize]string

		for {
			// read stdin, close the result channel if solver terminate
			line, isPrefix, err = input.ReadLine()
			switch {
			case isPrefix:
				panic("solver sends a big message")
			case err == io.EOF:
				log.Println("solver terminated")
				close(result)
				return
			case err != io.EOF && err != nil:
				panic("cannot read from stdin; err= " + err.Error())
			case len(line) != MaxCodeSize:
				panic("solver send bad input: " + string(line))
			}

			// convert to easy-to-read format
			for i := 0; i < MaxCodeSize; i++ {
				guesses[i] = string(line[i])
			}
			log.Printf("<:%s%s%s%s\\n\n", guesses[0], guesses[1], guesses[2], guesses[3])

			// queue up sovler output
			result <- guesses
		}
	}()

	return result
}

// expect to be give the commands to run solver in the args
func getSolver() *exec.Cmd {
	full := flag.Args()
	return exec.Command(full[0], full[1:]...)
}

func writeToSolver(solverStdin io.Writer, turn, colorAndPos, colorNoPos int, state string) {
	log.Printf(">:%d %d %d %s\\n\n", turn, colorAndPos, colorNoPos, state)
	fmt.Fprintf(solverStdin, "%d %d %d %s\n", turn, colorAndPos, colorNoPos, state)
}
