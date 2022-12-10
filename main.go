package main

import (
	"os"

	"github.com/kijimaD/gar/cmd"
	"github.com/kijimaD/gar/strutil"
)

func main() {
	c := cmd.New(os.Stdout, strutil.GetPrompt())
	err := c.Execute(os.Args)

	if err != nil {
		panic(err)
	}
}
