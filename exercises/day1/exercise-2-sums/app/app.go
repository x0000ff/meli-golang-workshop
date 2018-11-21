package app

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

func Run(args []string, stdout io.Writer) error {

	if len(args) < 2 {
		//printHelp(stdout)
		return errors.New("not enough arguments")
	}

	rawInput := args[1]
	input, err := strconv.Atoi(rawInput)
	if err != nil {
		return fmt.Errorf("incorrect input: \"%v\", expected integer", rawInput)
	}

	result := sum(input)
	_, err = fmt.Fprintln(stdout, result)
	return err
}

func GetHelp(filename string) string {

	return fmt.Sprintf(`Usage:
$ go run %s <MAX>
$ go run %s 31

Only integer values allowed`, filename, filename)

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
