package main

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/git-hook/git-ledger/src/git-ledger"
)

func find(c *cli.Context) error {
	slug := c.Args().First()

	record, err := ledger.GetBySlug(slug)
	if err != nil {
		// TODO: print a helpful message to stderr
		fmt.Println(err)
		// panic(err)
	}
	fmt.Println(record.Path)

	return nil
}
