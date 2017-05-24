package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func rm(c *cli.Context) error {
	fmt.Println("I am in rm!")
	return nil
}
