package main

import (
	"flag"
	"gar/gh"
)

func main() {
	var (
		n = flag.Int("n", -1, "PR number")
	)
	flag.Parse()
	if *n == -1 {
		panic("need PR number")
	}

	g, err := gh.New(*n)

	if err != nil {
		panic(err)
	}

	c := gh.NewClient(g, g.PR)
	c.GetCommits()
	c.ParseCommit()
	c.SendReply()
}
