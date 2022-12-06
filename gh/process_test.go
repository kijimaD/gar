package gh

import (
	"bytes"
	"os"
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
	s := NewClient(cl, os.Stdout)
	s.Commits = commits

	s.ParseCommit()

	expect := []Reply{
		{
			ReplyID:   int64(1037682054),
			GitHash:   sha0,
			CommitMsg: message0,
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

	s := NewClient(cl, os.Stdout)

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

func TestFetchComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	body := "original comment"
	cl.EXPECT().GetComment(gomock.Any()).Times(1).Return(&github.PullRequestComment{
		Body: &body,
	})

	s := NewClient(cl, os.Stdout)
	s.Replys = []Reply{
		{
			ReplyID:   int64(1037682054),
			GitHash:   "1111111",
			CommitMsg: "commit msg",
		},
	}
	s.FetchComment()

	for _, r := range s.Replys {
		assert.Equal(t, body, r.OriginalComment)
	}
}

func TestDisplay(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	t.Run("when exists reply", func(t *testing.T) {
		buffer := bytes.Buffer{}
		s := NewClient(cl, &buffer)

		s.Replys = []Reply{
			{
				ReplyID:   int64(1037682054),
				GitHash:   "1111111",
				CommitMsg: "try to fix problem",
			},
			{
				ReplyID:   int64(1037699999),
				GitHash:   "1122334",
				CommitMsg: "refactor long comment",
			},
			{
				ReplyID:   int64(1037699999),
				GitHash:   "1122334",
				CommitMsg: "typo",
			},
		}
		s.Display()

		got := buffer.String()
		expect := `The execution of this command will result in the following.
+-----+-------------------+-------------------+------+
| IDX |      COMMIT       |  LINKED COMMENT   | SEND |
+-----+-------------------+-------------------+------+
|  00 | 1111111 try to fi | this is original  | yes  |
|  01 | 1122334 refactor  | this is original  | yes  |
|  02 | 1122334 typo      | this is original  | yes  |
+-----+-------------------+-------------------+------+
`
		assert.Equal(t, expect, got)
	})

	t.Run("when not exists reply", func(t *testing.T) {
		buffer := bytes.Buffer{}
		s := NewClient(cl, &buffer)
		s.Replys = []Reply{}
		s.Display()

		got := buffer.String()
		expect := `The execution of this command will result in the following.
Not found reply target!
`
		assert.Equal(t, expect, got)
	})
}
