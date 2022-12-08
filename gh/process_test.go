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

func TestFetchCommentById(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	body := "original comment"
	cl.EXPECT().GetCommentByID(gomock.Any()).Times(1).Return(&github.PullRequestComment{
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
	s.FetchCommentById()

	for _, r := range s.Replys {
		assert.Equal(t, body, r.OriginalComment)
	}
}

func TestFetchPRComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	body0 := "comment0"
	body1 := "comment1"
	cl.EXPECT().GetCommentList().Times(1).Return([]*github.PullRequestComment{
		{
			Body: &body0,
		},
		{
			Body: &body1,
		},
	})

	s := NewClient(cl, os.Stdout)
	s.FetchPRComment()
	assert.Equal(t, "comment0 comment1", s.PRComment)
}

func TestValidate(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	buffer := bytes.Buffer{}
	s := NewClient(cl, &buffer)

	t.Run("original comment check", func(t *testing.T) {
		s.PRComment = ""
		s.Replys = []Reply{
			{
				OriginalComment: "comment",
				GitHash:         "111111",
			},
			{
				OriginalComment: "",
				GitHash:         "222222",
			},
		}

		s.Validate()

		assert.Equal(t, true, s.Replys[0].IsValid)
		assert.Equal(t, false, s.Replys[1].IsValid)
	})
	t.Run("duplicate comment check", func(t *testing.T) {
		s.PRComment = "111111 222222"
		s.Replys = []Reply{
			{
				OriginalComment: "comment",
				GitHash:         "111111",
			},
			{
				OriginalComment: "comment",
				GitHash:         "222222",
			},
			{
				OriginalComment: "comment",
				GitHash:         "333333",
			},
		}

		s.Validate()

		assert.Equal(t, false, s.Replys[0].IsValid)
		assert.Equal(t, false, s.Replys[1].IsValid)
		assert.Equal(t, true, s.Replys[2].IsValid)
	})
}

func TestDisplay(t *testing.T) {
	ctrl := gomock.NewController(t)
	cl := NewMockclientI(ctrl)

	t.Run("when exists reply", func(t *testing.T) {
		buffer := bytes.Buffer{}
		s := NewClient(cl, &buffer)

		s.Replys = []Reply{
			{
				ReplyID:         int64(1037682054),
				GitHash:         "1111111",
				CommitMsg:       "try to fix problem",
				OriginalComment: "original comment0",
			},
			{
				ReplyID:         int64(1037699999),
				GitHash:         "1122334",
				CommitMsg:       "refactor long comment",
				OriginalComment: "original comment1",
			},
			{
				ReplyID:         int64(1037699999),
				GitHash:         "1122334",
				CommitMsg:       "typo",
				OriginalComment: "original comment2",
			},
		}
		s.Display()

		got := buffer.String()
		expect := `The execution of this command will result in the following.
+-----+-------------------+-------------------+------+
| IDX |      COMMIT       |  LINKED COMMENT   | SEND |
+-----+-------------------+-------------------+------+
|  00 | 1111111 try to fi | original comment0 | no   |
|  01 | 1122334 refactor  | original comment1 | no   |
|  02 | 1122334 typo      | original comment2 | no   |
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
