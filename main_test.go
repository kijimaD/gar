//go:build gh

package main

import (
	"gar/gh"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	g, err := gh.New(1)

	if err != nil {
		panic(err)
	}

	c := gh.NewClient(g, os.Stdout)
	c.GetCommits()
	c.ParseCommit()
	c.Display()
	c.FetchCommentById()
	c.FetchPRComment()
	c.SendReply()
}
