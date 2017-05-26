package main

import (
	"fmt"
	"os"

	"github.com/git-hook/git-ledger"
	"github.com/urfave/cli"
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
