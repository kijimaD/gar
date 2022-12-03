package gh

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
)

func TestPRGitHashs(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	number := 1
	hRef := "dev"
	head := github.PullRequestBranch{
		Ref: &hRef,
	}
	bRef := "main"
	base := github.PullRequestBranch{
		Ref: &bRef,
	}

	cl.EXPECT().PRDetail().AnyTimes().Return(&github.PullRequest{
		Number: &number,
		Head:   &head,
		Base:   &base,
	})

	s := &CallClient{
		API: cl,
	}
	s.PRGitHashs()
}
