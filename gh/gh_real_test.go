//go:build gh

package gh

import (
	"testing"
)

// 実際にGitHubにアクセスするテスト

func TestListReal(t *testing.T) {
	gh, err := New()

	if err != nil {
		t.Error(err)
	}

	gh.List()
}

func TestReplyReal(t *testing.T) {
	gh, err := New()

	if err != nil {
		t.Error(err)
	}

	gh.Reply()
}
