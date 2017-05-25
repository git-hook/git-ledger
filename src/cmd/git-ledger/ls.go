package main

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/git-hook/git-ledger/src/git-ledger"
)

func ls(c *cli.Context) error {
	records, err := ledger.GetRecords()
	if err != nil {
		// TODO: stderr time
		// TODO exit code
		panic(err)
	}

	for _, rec := range records {
		fmt.Println(fmt.Sprintf("%s: %s", rec.Slug, rec.Path))
	}
	// TODO: handle no size, print "none" to stderr? if git does something similar
	return nil
}
