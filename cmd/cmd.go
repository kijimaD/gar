//go:generate mockgen -source=cmd.go -destination=cmd_mock.go -package=cmd

package cmd

import (
	"errors"
	"flag"
	"io"
	"strconv"

	"github.com/kijimaD/gar/gh"
)

type CLI struct {
	Stdout io.Writer
	Runner Runner
}

type Runner interface {
	Run() (int, string, error)
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

func New(stdout io.Writer, runner Runner) *CLI {
	return &CLI{
		Stdout: stdout,
		Runner: runner,
	}
}

func (cli *CLI) Execute(args []string) error {
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
	c.Validate()
	c.Display()

	_, confirm, err := cli.Runner.Run()
	if err != nil {
		panic(err)
	}

	if confirm == "yes" || f {
		err := c.SendReply()
		if err != nil {
			return err
		}
	}

	return nil
}
