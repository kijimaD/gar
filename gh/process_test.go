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

	message := "this is commit message"
	commit := github.Commit{
		Message: &message,
	}
	rc := github.RepositoryCommit{
		Commit: &commit,
	}
	commits := []*github.RepositoryCommit{&rc}
	s := &CallClient{
		API:     cl,
		Commits: commits,
	}

	s.ParseCommit()

	// コミットメッセージをパースできてることを確認する
	// 返り値の[]Replyが正しいかどうか?

	expect := []Reply{
		{
			ReplyID: 1,
			GitHash: "a3d",
		},
	}

	assert.Equal(t, expect, s.Replys)
}

func TestParseMsg(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	s := &CallClient{
		API: cl,
	}

	expect := "1037682054"

	result := s.ParseMsg(`feat: this is test

https://github.com/kijimaD/gar/pull/1#discussion_r1037682054`)

	assert.Equal(t, expect, result)
}
