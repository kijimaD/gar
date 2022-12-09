package cmd

import (
	"errors"
	"flag"
	"io"
	"strconv"

	"github.com/kijimaD/gar/gh"
	"github.com/kijimaD/gar/strutil"
)

type CLI struct {
	Stdout io.Writer
	// Stderr io.Writer
	// Stdin  io.Reader
}

var (
	f bool
)

var (
	PRNumNotExistError = errors.New("require PR number argument")
)

func init() {
	flag.BoolVar(&f, "force", false, "send reply")
}

func New(stdout io.Writer) *CLI {
	return &CLI{
		Stdout: stdout,
	}
}

func (cli *CLI) Run(args []string) error {
	flag.Parse()

	if len(args) <= 1 {
		return PRNumNotExistError
	}

	n, err := strconv.Atoi(args[1])
	if err != nil {
		return err
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
	strutil.Ask()
	if f {
		c.SendReply()
	}

	return nil
}
