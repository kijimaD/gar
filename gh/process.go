package gh

import (
	"fmt"
	"log"
	"strconv"

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
	var replys []Reply

	for _, commit := range c.Commits {
		id := c.parseMsg(*commit.Commit.Message)

		reply := Reply{
			ReplyID: id,
			GitHash: *commit.SHA,
		}

		replys = append(replys, reply)
	}

	c.Replys = replys
}

func (c *CallClient) parseMsg(string) int64 {
	result, err := strconv.ParseInt("1037682054", 10, 64)
	if err != nil {
		log.Println("can't parse")
	}
	return result
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
