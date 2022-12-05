//go:generate mockgen -source=gh.go -destination=gh_mock.go -package=gh

package gh

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

var (
	NotFoundGitRemoteError = errors.New("not found remote URL...")
)

type clientI interface {
	SendReply(Reply)
	PRDetail() *github.PullRequest
	PRCommits() []*github.RepositoryCommit
}

type GitHub struct {
	Client *github.Client
	PR     PR
}

type PR struct {
	User   string
	Repo   string
	Number int
}

func New(PRnumber int) (*GitHub, error) {
	token := os.Getenv("GH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	pr, err := getGitInfo()
	if err != nil {
		panic(err)
	}
	pr.Number = PRnumber

	return &GitHub{
		Client: client,
		PR:     *pr,
	}, nil
}

func getGitInfo() (*PR, error) {
	out, err := exec.Command("git", "config", "--get", "remote.origin.url").Output()
	if err != nil {
		return &PR{}, NotFoundGitRemoteError
	}

	raw := string(out)

	r := regexp.MustCompile(`git@github.com:(\S+)/(\S+).git`)
	result := r.FindAllStringSubmatch(raw, -1)

	if len(result) == 0 {
		return &PR{}, fmt.Errorf("%s", err)
	}

	user := result[0][1]
	repo := result[0][2]

	return &PR{
		User: user,
		Repo: repo,
	}, nil
}

func (gh *GitHub) SendReply(r Reply) {
	ctx := context.Background()
	msg := fmt.Sprintf("check 👉 %s", r.GitHash)
	_, _, err := gh.Client.PullRequests.CreateCommentInReplyTo(ctx, gh.PR.User, gh.PR.Repo, gh.PR.Number, msg, r.ReplyID)

	if err != nil {
		panic(err)
	}
}

func (gh *GitHub) PRDetail() *github.PullRequest {
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

func (gh *GitHub) PRCommits() []*github.RepositoryCommit {
	ctx := context.Background()
	commits, _, err := gh.Client.PullRequests.ListCommits(ctx, gh.PR.User, gh.PR.Repo, gh.PR.Number, nil)

	if err != nil {
		panic(err)
	}

	// for _, c := range commits {
	// 	fmt.Printf("%v\n", c)
	// 	fmt.Printf("%v\n", *c.Commit.Message)
	// 	fmt.Printf("%v\n", *c.SHA)
	// }

	return commits
}
