package main

import (
	"flag"
	"gar/gh"
	"os"
)

func main() {
	var (
		n = flag.Int("n", -1, "PR number")
	)
	flag.Parse()
	if *n == -1 {
		panic("error: need PR number option. -n {number}")
	}

	g, err := gh.New(*n)

	if err != nil {
		panic(err)
	}

	c := gh.NewClient(g, os.Stdout)
	c.GetCommits()
	c.ParseCommit()
	c.SendReply()
}
