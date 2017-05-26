package main

import (
	"os"
	"fmt"

	"github.com/urfave/cli"
	"github.com/git-hook/git-ledger"
)

// Find path associated with specified path or slug from the ledger.
func find(c *cli.Context) error {
	input := c.Args().First()

	record, err := ledger.GetBySlug(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err
	}
	fmt.Println(record.Path)

	return nil
}
