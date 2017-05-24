package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/urfave/cli"
	"github.com/git-hook/git-ledger/src/git-ledger"
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
	fmt.Println("I am in add!")
	// Default: add cwd to toml
	project := c.Args().First()

	// TODO: clean up this declaration
	var record ledger.Record
	record.Path = project
	record.Slug = getSlug(project)

	fmt.Println("record is ", record)
	fmt.Println("path is ", ledger.Path())

	record.RemoveFromLedger()
	record.WriteToLedger()

	return nil
}
