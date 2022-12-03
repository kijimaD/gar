//go:build gh

package gh

import "testing"

// 実際にGitHubにアクセスするテスト

func TestListReal(t *testing.T) {
	gh, err := New()

	if err != nil {
		t.Errorf("%s", err)
	}

	gh.List()
}
