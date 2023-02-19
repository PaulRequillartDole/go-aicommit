package main

import (
	"fmt"
	"github.com/PaulRequillartDole/go-aicommit/internal/command"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("AiCommit version %s 🤖\n", cCtx.App.Version)
	}
	app := &cli.App{
		Name:    "AiCommit",
		Usage:   "Generate commit message with AI 🤖",
		Version: "v0.1",
		Action: func(*cli.Context) error {
			command.Execute()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
