package main

import (
	"os"

	"github.com/owsky/game-libraries-crosschecker/checker"
	"github.com/owsky/game-libraries-crosschecker/config"
)

func main() {
	args := os.Args[1:]
	config.SetEnvironmentVariables(args)
	checker.Crosscheck()
}
