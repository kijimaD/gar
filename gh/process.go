package gh

import (
	"fmt"
	"gar/strutil"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/go-github/v48/github"
	"github.com/olekukonko/tablewriter"
)

type CallClient struct {
	API       clientI
	Writer    io.Writer
	Commits   []*github.RepositoryCommit
	Replys    []Reply
	PRComment string
}

func NewClient(api clientI, writer io.Writer) *CallClient {
	return &CallClient{
		API:    api,
		Writer: writer,
	}
}

type Reply struct {
	ReplyID         int64
	GitHash         string
	CommitMsg       string
	OriginalComment string
	IsValid         bool
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
			ReplyID:   id,
			GitHash:   *commit.SHA,
			CommitMsg: *commit.Commit.Message,
		}

		replys = append(replys, reply)
	}

	c.Replys = replys
}

func (c *CallClient) parseMsg(s string) (int64, error) {
	user := c.API.GetPR().User
	repo := c.API.GetPR().Repo
	num := c.API.GetPR().Number
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

func (c *CallClient) FetchCommentById() {
	for i, r := range c.Replys {
		comment := c.API.GetCommentByID(r.ReplyID)

		c.Replys[i].OriginalComment = *comment.Body
	}
}

func (c *CallClient) FetchPRComment() {
	comments := c.API.GetCommentList()
	var commentsStr []string
	for _, com := range comments {
		commentsStr = append(commentsStr, *com.Body)
	}

	c.PRComment = strings.Join(commentsStr, " ")
}

func (c *CallClient) Validate() {
	for i, r := range c.Replys {
		// 返信先があるかどうか
		valid1 := false
		if r.OriginalComment == "" {
			valid1 = false
		} else {
			valid1 = true
		}

		// 重複してるかどうか
		valid2 := false
		if strings.Contains(c.PRComment, r.GitHash) {
			valid2 = false
		} else {
			valid2 = true
		}

		c.Replys[i].IsValid = valid1 && valid2
	}
}

func (c *CallClient) Display() {
	fmt.Fprintln(c.Writer, "The execution of this command will result in the following.")

	if len(c.Replys) == 0 {
		fmt.Fprintf(c.Writer, "Not found reply target!\n")
	} else {
		data := [][]string{}

		for i, r := range c.Replys {
			data = append(data, []string{
				fmt.Sprintf("%02d", i),
				fmt.Sprintf(
					"%s %s",
					r.GitHash[0:7],
					strutil.Substring(r.CommitMsg, 0, 9)),
				strutil.Substring(r.OriginalComment, 0, 17),
				strutil.YorN(r.IsValid),
			})
		}

		table := tablewriter.NewWriter(c.Writer)
		table.SetHeader([]string{"IDX", "COMMIT", "LINKED COMMENT", "SEND"})
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	}
}

func (c *CallClient) SendReply() {
	for _, r := range c.Replys {
		c.API.SendReply(Reply{
			ReplyID: r.ReplyID,
			GitHash: r.GitHash,
		})
	}
}
