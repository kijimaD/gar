package gh

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/stretchr/testify/assert"
)

func TestShowHash(t *testing.T) {
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
	s.showHash()
}

func TestParseCommit(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	message := `this is commit message
https://github.com/kijimaD/gar/pull/1#discussion_r1037682054
`
	sha := "369a79d9028e378b2f4ad3e566df061583656617"
	commit := github.Commit{
		Message: &message,
	}
	rc := github.RepositoryCommit{
		Commit: &commit,
		SHA:    &sha,
	}
	commits := []*github.RepositoryCommit{&rc}
	s := &CallClient{
		API:     cl,
		Commits: commits,
	}

	s.ParseCommit()

	expect := []Reply{
		{
			ReplyID: int64(1037682054),
			GitHash: sha,
		},
	}

	// Replysに値がセットされている
	assert.Equal(t, expect, s.Replys)
}

func TestParseMsg(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	s := &CallClient{
		API: cl,
	}

	expect := int64(1037682054)

	result, err := s.parseMsg(`feat: this is test

https://github.com/kijimaD/gar/pull/1#discussion_r1037682054`)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expect, result)
}
