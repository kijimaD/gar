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
	Reply()
	PRDetail()
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

func (gh *Gh) PRDetail() {
	ctx := context.Background()
	_, _, err := gh.Client.PullRequests.Get(ctx, "kijimaD", "gar", 1)

	if err != nil {
		panic(err)
	}
}

type CallClient struct {
	API clientI
}

func (sync *CallClient) process() {
	sync.API.PRDetail()
	fmt.Println("execute!")

}
