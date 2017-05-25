package main

// TODO: test
// TODO: doxygenize

import (
	"fmt"
	"regexp"
	"strings"
	"path/filepath"

	"github.com/urfave/cli"
	"github.com/codeskyblue/go-sh"
	"github.com/git-hook/git-ledger/src/git-ledger"
)

func getRemote(dir string) string {
	session := sh.NewSession()
	session.SetDir(dir)
	remoteRaw, err := session.Command("git", "remote").Command("head", "-n", "1").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(fmt.Sprintf("%s", remoteRaw))
}

func getSlug(dir string) string {
	session := sh.NewSession()
	session.SetDir(dir)
	out, err := session.Command("git", "remote", "get-url", getRemote(dir)).Output()
	output := strings.TrimSpace(fmt.Sprintf("%s", out))
	reg := regexp.MustCompile(`[^/:]*/[^/:]*$`)
	res := reg.FindStringSubmatch(output)
	if err != nil {
		panic(err)
	}
	return res[0]
}

func add(c *cli.Context) error {
	// TODO: print message to user
	fmt.Println("I am in add!")
	// Default: add cwd to toml
	project := c.Args().First()

	// TODO: ensure project contains a '.git' directory

	// TODO: clean up this declaration
	var record ledger.Record
	record.Path, _ = filepath.Abs(project)
	record.Slug = getSlug(project)

	fmt.Println("record is ", record)
	fmt.Println("path is ", ledger.Path())

	record.RemoveFromLedger()
	record.AddToLedger()

	return nil
}
