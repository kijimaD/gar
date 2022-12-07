package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("valid PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}
		c := New(&buffer)
		err := c.Run([]string{"any", "1"})

		if err != nil {
			t.Error(err)
		}
	})
	t.Run("not exist PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}
		c := New(&buffer)
		err := c.Run([]string{"any"})

		if err == nil {
			t.Error("expect: fail")
		}
		assert.Equal(t, PRNumNotExistError, err)
	})
	t.Run("invalid PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}
		c := New(&buffer)
		err := c.Run([]string{"any", "あ"})

		if err == nil {
			t.Error("expect: fail")
		}
	})
}
