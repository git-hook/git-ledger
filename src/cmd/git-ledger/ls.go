package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func ls(c *cli.Context) error {
	fmt.Println("I am in ls!")
	return nil
}
