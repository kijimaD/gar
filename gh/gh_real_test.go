//go:build gh

// 実際にGitHubにアクセスするテスト

package gh

import (
	"testing"
)

// レビューコメントに返信する
func TestReplyReal(t *testing.T) {
	gh, err := New(1)

	if err != nil {
		t.Error(err)
	}

	gh.SendReply(
		Reply{
			ReplyID: int64(1037682054),
			GitHash: "111111",
		},
	)
}

func TestPRCommitsReal(t *testing.T) {
	gh, err := New(1)

	if err != nil {
		t.Error(err)
	}

	gh.PRCommits()
}

func TestGetComment(t *testing.T) {
	gh, err := New(1)

	if err != nil {
		t.Error(err)
	}

	gh.GetComment(1037682054)
}

func TestGetCommentList(t *testing.T) {
	gh, err := New(1)

	if err != nil {
		t.Error(err)
	}

	gh.GetCommentList()
}
