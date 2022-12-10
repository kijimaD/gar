package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	t.Run("valid PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}
		c := New(&buffer)
		err := c.Execute([]string{"any", "1"})

		if err != nil {
			t.Error(err)
		}
	})
	t.Run("not exist PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}
		c := New(&buffer)
		err := c.Execute([]string{"any"})

		if err == nil {
			t.Error("expect: fail")
		}
		assert.Equal(t, PRNumNotExistError, err)
	})
	t.Run("invalid PR number arg", func(t *testing.T) {
		buffer := bytes.Buffer{}
		c := New(&buffer)
		err := c.Execute([]string{"any", "あ"})

		if err == nil {
			t.Error("expect: fail")
		}
	})
}
