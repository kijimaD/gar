package gh

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type Gh struct {
	client *github.Client
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
		client: client,
	}, nil
}

func (g *Gh) Client() *github.Client {
	return g.client
}

func (g *Gh) SetClient(client *github.Client) {
	g.client = client
}

func Main() {
	gh, err := New()

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	repos, _, err := gh.Client().Repositories.List(ctx, "kijimad", nil)

	fmt.Printf("%#v", repos)

	if err != nil {
		panic(err)
	}
}
