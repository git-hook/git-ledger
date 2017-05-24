package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
	// "github.com/BurntSushi/toml"
	"github.com/ericcrosson/git-ledger/src/git-ledger"
	"github.com/codeskyblue/go-sh"
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
	out, err := session.Command("git", "remote", "get-url", getRemote(dir)).Command("sed", "-r", "s#[^/:]*/##").Output()
	output := strings.TrimSpace(fmt.Sprintf("%s", out))
	if err != nil {
		panic(err)
	}
	return output
}

func add(c *cli.Context) error {
	fmt.Println("I am in add!")
	// Default: add cwd to toml
	project := c.Args().First()

	// TODO: clean up this declaration
	var record ledger.Record
	record.Path = project
	record.Slug = getSlug(project)

	fmt.Println("record is ", record)
	fmt.Println("path is ", ledger.Path())

	return nil
}
