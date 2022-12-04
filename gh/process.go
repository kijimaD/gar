package gh

import (
	"fmt"

	"github.com/google/go-github/v48/github"
)

type CallClient struct {
	PR      PR
	API     clientI
	Commits []*github.RepositoryCommit
	Replys  []Reply
}

func (c *CallClient) showHash() {
	pr := c.API.PRDetail()
	curBranch := *pr.Head.Ref
	baseBranch := *pr.Base.Ref
	fmt.Println(curBranch, baseBranch)
}

type Reply struct {
	ReplyID int64  // 1037682054
	GitHash string // 90a142
}

func (c *CallClient) GetCommits() {
	c.Commits = c.API.PRCommits()
}

func (c *CallClient) ParseCommit() {
	c.Replys = []Reply{
		{
			ReplyID: 1,
			GitHash: "a3d",
		},
	}
}

func (c *CallClient) ParseMsg(string) string {
	return "1037682054"
}

func (c *CallClient) SendReply() {
	c.API.Reply(Reply{
		ReplyID: 1,
		GitHash: "a3d",
	})
}

func main() {
	gh, err := New()

	if err != nil {
		panic(err)
	}

	c := &CallClient{
		API: gh,
	}
	c.GetCommits()
	c.ParseCommit()
	c.SendReply()
}
