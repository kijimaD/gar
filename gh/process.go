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

func NewClient(api clientI, pr PR) *CallClient {
	return &CallClient{
		API: api,
		PR:  pr,
	}
}

type Reply struct {
	ReplyID int64
	GitHash string
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
	user := c.PR.User
	repo := c.PR.Repo
	num := c.PR.Number
	regex := fmt.Sprintf(`https://github.com/%s/%s/pull/%d#discussion_r(\d+)`, user, repo, num)
	r := regexp.MustCompile(regex)

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
	for _, r := range c.Replys {
		c.API.SendReply(Reply{
			ReplyID: r.ReplyID,
			GitHash: r.GitHash,
		})
	}
}
