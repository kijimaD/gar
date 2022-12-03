//go:generate mockgen -source=gh.go -destination=gh_mock.go -package=gh

package gh

import (
	"context"
	"os"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type clientI interface {
	Reply()
	PRDetail() *github.PullRequest
	PRCommits() []*github.RepositoryCommit
}

type Gh struct {
	Client *github.Client
}

func New() (*Gh, error) {
	token := os.Getenv("GH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return &Gh{
		Client: client,
	}, nil
}

func (gh *Gh) Reply() {
	ctx := context.Background()
	_, _, err := gh.Client.PullRequests.CreateCommentInReplyTo(ctx, "kijimaD", "gar", 1, "this is test by API", 1037682054)

	if err != nil {
		panic(err)
	}
}

func (gh *Gh) PRDetail() *github.PullRequest {
	ctx := context.Background()
	pr, _, err := gh.Client.PullRequests.Get(ctx, "kijimaD", "gar", 1)

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
	commits, _, err := gh.Client.PullRequests.ListCommits(ctx, "kijimaD", "gar", 1, nil)

	if err != nil {
		panic(err)
	}

	// for _, c := range commits {
	// 	fmt.Printf("%v\n", c)
	// 	fmt.Printf("%v\n", *c.Commit.Message)
	// }

	return commits
}
