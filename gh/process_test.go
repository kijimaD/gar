package gh

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/stretchr/testify/assert"
)

func TestParseCommit(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)
	cl.EXPECT().GetPR().Times(6).Return(PR{
		User:   "kijimaD",
		Repo:   "gar",
		Number: 1,
	})

	message0 := `this is commit message

https://github.com/kijimaD/gar/pull/1#discussion_r1037682054`
	sha0 := "369a79d9028e378b2f4ad3e566df061583656617"
	commit0 := github.Commit{
		Message: &message0,
	}
	rc0 := github.RepositoryCommit{
		Commit: &commit0,
		SHA:    &sha0,
	}

	message1 := "not contain reply URL"
	sha1 := "369a79d9028e378b2f4ad3e566df061583656617"
	commit1 := github.Commit{
		Message: &message1,
	}
	rc1 := github.RepositoryCommit{
		Commit: &commit1,
		SHA:    &sha1,
	}

	commits := []*github.RepositoryCommit{&rc0, &rc1}
	s := NewClient(cl)
	s.Commits = commits

	s.ParseCommit()

	expect := []Reply{
		{
			ReplyID: int64(1037682054),
			GitHash: sha0,
		},
	}

	// Replysに値がセットされている
	assert.Equal(t, expect, s.Replys)
}

func TestParseMsg(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)
	cl.EXPECT().GetPR().Times(9).Return(PR{
		User:   "kijimaD",
		Repo:   "gar",
		Number: 1,
	})

	s := NewClient(cl)

	expect := int64(1037682054)

	t.Run("パースできる", func(t *testing.T) {
		msg := `feat: this is test

https://github.com/kijimaD/gar/pull/1#discussion_r1037682054`

		result, err := s.parseMsg(msg)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, expect, result)
	})

	t.Run("PRが異なる場合マッチしない", func(t *testing.T) {
		msg := `feat: this is test

https://github.com/kijimaD/gar/pull/2#discussion_r1037682054`

		result, err := s.parseMsg(msg)

		if err != nil {
			t.Error(err)
		}

		if result != -1 {
			t.Error(err)
		}
	})

	t.Run("含まれない場合マッチしない", func(t *testing.T) {
		msg := `feat: this is test`

		result, err := s.parseMsg(msg)

		if err != nil {
			t.Error(err)
		}

		if result != -1 {
			t.Error(err)
		}
	})
}
