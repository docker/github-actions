package main

import (
	"fmt"
	"os"

	commandLine "github.com/urfave/cli/v2"
)

func main() {
	app := &commandLine.App{
		Name:  "docker github actions",
		Usage: "Used in GitHub Actions to run Docker workflows",
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
