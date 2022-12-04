package gh

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGitInfo(t *testing.T) {
	t.Run("git管理下で実行して取得できる", func(t *testing.T) {
		expect := &PR{
			User: "kijimaD",
			Repo: "gar",
		}

		result, err := getGitInfo()
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expect, result)
	})

	t.Run("git管理下でないとエラーになる", func(t *testing.T) {
		expect := &PR{}

		curPath, err := os.Getwd()
		if err != nil {
			t.Error(err)
		}
		err = os.Chdir("/")
		if err != nil {
			t.Error(err)
		}

		result, err := getGitInfo()

		if err == nil {
			t.Error(err)
		}
		assert.Equal(t, NotFoundGitRemoteError, err)
		assert.Equal(t, expect, result)

		// 戻す
		err = os.Chdir(curPath)
		if err != nil {
			t.Error(err)
		}
	})
}
