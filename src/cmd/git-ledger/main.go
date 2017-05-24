package main

import (
	"os"
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/jmalloc/grit/src/grit/pathutil"
	"github.com/urfave/cli"
)

var VERSION = semver.MustParse("0.0.0")

func main () {

	app := cli.NewApp()
	homeDir, _ := pathutil.HomeDir()

	app.Name = "grip"
	app.Usage = "Index your Git clones."
	app.Version = VERSION.String()

	app.Commands = []cli.Command {
		{
			Name:   "add",
			Usage:  "Start tracking an existing repository.",
			ArgsUsage: "[<path>]",
			Action: add,
		},
		{
			Name:   "find",
			Usage:  "Print the location of a tracked repository.",
			ArgsUsage: "[<path>]",
			Action: find,
		},
		{
			Name:   "ls",
			Usage:  "Print all tracked repositories",
			Action: find,
		},
		{
			Name:   "rm",
			Usage:  "Stop tracking an existing repository.",
			ArgsUsage: "[<path>]",
			Action: rm,
		},
	}

	fmt.Println("Made it here with homeDir ", homeDir)

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
