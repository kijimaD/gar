//go:build gh

// 実際にGitHubにアクセスするテスト

package gh

import (
	"testing"
)

// レビューコメントに返信する
func TestReplyReal(t *testing.T) {
	gh, err := New()

	if err != nil {
		t.Error(err)
	}

	gh.Reply()
}

// PRの詳細情報を取得する
func TestPRDetailReal(t *testing.T) {
	gh, err := New()

	if err != nil {
		t.Error(err)
	}

	gh.PRDetail()
}

func TestPRCommitsReal(t *testing.T) {
	gh, err := New()

	if err != nil {
		t.Error(err)
	}

	gh.PRCommits()
}
