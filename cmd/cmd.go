package cmd

import (
	"errors"
	"flag"
	"gar/gh"
	"io"
)

type CLI struct {
	Stdout io.Writer
	// Stderr io.Writer
	// Stdin  io.Reader
}

var (
	n int
	f bool
)

var (
	PRNumNotExistError = errors.New("error: need PR number option. -n {number}")
)

func init() {
	flag.IntVar(&n, "n", -1, "PR number")
	flag.BoolVar(&f, "false", false, "send reply")
}

func New(stdout io.Writer) *CLI {
	return &CLI{
		Stdout: stdout,
	}
}

func (cli *CLI) Run() error {
	flag.Parse()
	if n == -1 {
		return PRNumNotExistError
	}

	g, err := gh.New(n)

	if err != nil {
		panic(err)
	}

	c := gh.NewClient(g, cli.Stdout)
	c.GetCommits()
	c.ParseCommit()
	c.FetchCommentById()
	c.FetchPRComment()
	c.Display()
	if f {
		c.SendReply()
	}

	return nil
}
