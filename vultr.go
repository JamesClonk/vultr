package main

import (
	"os"

	"github.com/JamesClonk/vultr/cmd"
)

func main() {
	app := cmd.NewApp()
	app.RegisterCommands()
	app.Run(os.Args)
}
