package gh

import (
	"fmt"
	"regexp"
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
		id, err := c.parseMsg(*commit.Commit.Message)

		if err != nil {
			panic(err)
		}

		if id == -1 {
			continue
		}

		reply := Reply{
			ReplyID: id,
			GitHash: *commit.SHA,
		}

		replys = append(replys, reply)
	}

	c.Replys = replys
}

func (c *CallClient) parseMsg(s string) (int64, error) {
	r := regexp.MustCompile(`https://github.com/kijimaD/gar/pull/1#discussion_r(\d+)`)

	result := r.FindAllStringSubmatch(s, -1)

	if len(result) == 0 {
		return -1, nil
	}

	int64, err := strconv.ParseInt(result[0][1], 10, 64)
	if err != nil {
		return -1, err
	}
	return int64, nil
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
