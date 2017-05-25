package main

import (
	"os"
	"fmt"

	"github.com/urfave/cli"
	"github.com/git-hook/git-ledger"
)

// List all records in the git-ledger.
func ls(c *cli.Context) error {
	records, err := ledger.GetRecords()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
		return err
	}

	for _, rec := range records {
		fmt.Println(fmt.Sprintf("%s: %s", rec.Slug, rec.Path))
	}

	// Ledger is empty
	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Ledger is empty\n")
	}
	return nil
}
