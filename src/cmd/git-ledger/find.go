package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func find(c *cli.Context) error {
	fmt.Println("I am in find!")
	return nil
}
