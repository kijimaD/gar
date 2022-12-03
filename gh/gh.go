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
	List()
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

func (gh *Gh) List() {
	ctx := context.Background()
	repos, _, err := gh.Client.Repositories.List(ctx, "kijimad", nil)

	fmt.Printf("%#v", repos)

	if err != nil {
		panic(err)
	}
}

type CallClient struct {
	API clientI
}

func (sync *CallClient) run() {
	fmt.Println("run")
	sync.API.List()
}
