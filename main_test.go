package main

import (
	"gar/gh"
	"testing"
)

func TestMain(t *testing.T) {
	g, err := gh.New(1)

	if err != nil {
		panic(err)
	}

	c := gh.CallClient{
		API: g,
	}
	c.GetCommits()
	c.ParseCommit()
	c.SendReply()
}
