package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
		return
	}

	rawInput := os.Args[1]
	input, err := strconv.Atoi(rawInput)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Incorrect input: \"%v\", expected integer\n", rawInput)
		os.Exit(1)
		return
	}

	result := sum(input)
	_, _ = fmt.Fprint(os.Stdout, result)
}

func printHelp() {
	filename := path.Base(os.Args[0])

	_, _ = fmt.Fprintf(os.Stdout, `Usage:
$ go run %s <MAX>
$ go run %s 31

Only integer values allowed
`, filename, filename)
}

func sum(max int) int {

	if max <= 0 {
		return 0
	}

	sum := 0
	for i := 1; i <= max; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}

	return sum
}
