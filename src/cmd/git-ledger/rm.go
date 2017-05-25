package main

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/git-hook/git-ledger/src/git-ledger"
)

// TODO document that this command:
// takes PATH

// TODO: set exit code (project-wide)
func rm(c *cli.Context) error {
	slug := c.Args().First()

	record, err := ledger.GetBySlug(slug)
	if err != nil {
		// TODO: print a helpful message to stderr
		fmt.Println(err)
		// panic(err)
	}

	err = record.RemoveFromLedger()
	if err != nil {
		// TODO: print a helpful message to stderr
		fmt.Println(err)
		// panic(err)
	}

	return nil
}
