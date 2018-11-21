package main

import (
	"fmt"
	"os"
	"path"
	"x0000ff/meli-golang-workshop/exercises/day1/exercise-2-sums/app"
)

func main() {
	err := app.Run(os.Args, os.Stdout)
	if err != nil {
		help := app.GetHelp(path.Base(os.Args[0]))
		_, _ = fmt.Fprintln(os.Stderr, err)
		_, _ = fmt.Fprintln(os.Stdout, help)
	}
}
