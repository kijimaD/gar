package gh

import (
	"fmt"

	"github.com/google/go-github/v48/github"
)

type CallClient struct {
	API clientI
}

func (c *CallClient) showHash() {
	pr := c.API.PRDetail()
	curBranch := *pr.Head.Ref
	baseBranch := *pr.Base.Ref
	fmt.Println(curBranch, baseBranch)
}

type Reply struct {
	ReplyID string // 1037682054
	GitHash string // 90a142
}

func (c *CallClient) ParseCommit([]*github.RepositoryCommit) []Reply {
	return []Reply{
		{
			ReplyID: "1",
			GitHash: "a3d",
		},
	}
}

func (c *CallClient) SendReply([]Reply) {
	c.API.Reply()
}

func (c *CallClient) ParseMsg(string) string {
	return "1037682054"
}
