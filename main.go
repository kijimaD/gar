package main

import (
	"os"

	"github.com/kijimaD/gar/cmd"
)

func main() {
	c := cmd.New(os.Stdout)
	err := c.Execute(os.Args)

	if err != nil {
		panic(err)
	}
}
