package main

import (
	"gar/cmd"
	"os"
)

func main() {
	c := cmd.New(os.Stdout)
	err := c.Run()

	if err != nil {
		panic(err)
	}
}
