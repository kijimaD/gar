//go:generate mockgen -source=gh.go -destination=gh_mock.go -package=gh

package gh

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type clientI interface {
	Reply(Reply)
	PRDetail() *github.PullRequest
	PRCommits() []*github.RepositoryCommit
}

type Gh struct {
	Client *github.Client
	PR     PR
}

type PR struct {
	User   string
	Repo   string
	Number int
}

func New() (*Gh, error) {
	token := os.Getenv("GH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	pr := PR{
		User:   "kijimaD",
		Repo:   "gar",
		Number: 1,
	}

	return &Gh{
		Client: client,
		PR:     pr,
	}, nil
}

func (gh *Gh) Reply(content Reply) {
	ctx := context.Background()
	_, _, err := gh.Client.PullRequests.CreateCommentInReplyTo(ctx, gh.PR.User, gh.PR.Repo, gh.PR.Number, content.GitHash, content.ReplyID)

	if err != nil {
		panic(err)
	}
}

func (gh *Gh) PRDetail() *github.PullRequest {
	ctx := context.Background()
	pr, _, err := gh.Client.PullRequests.Get(ctx, gh.PR.User, gh.PR.Repo, gh.PR.Number)

	if err != nil {
		panic(err)
	}

	// fmt.Println("pr================")
	// fmt.Printf("%v\n", pr)
	// fmt.Println("head================")
	// fmt.Printf("%v\n", *pr.Head.Ref) // PRのブランチ
	// fmt.Println("commits================")
	// fmt.Printf("%v\n", *pr.Base.Ref) // PRのベースブランチ

	return pr
}

func (gh *Gh) PRCommits() []*github.RepositoryCommit {
	ctx := context.Background()
	commits, _, err := gh.Client.PullRequests.ListCommits(ctx, gh.PR.User, gh.PR.Repo, gh.PR.Number, nil)

	if err != nil {
		panic(err)
	}

	for _, c := range commits {
		fmt.Printf("%v\n", c)
		fmt.Printf("%v\n", *c.Commit.Message)
		fmt.Printf("%v\n", *c.SHA)
	}

	return commits
}
